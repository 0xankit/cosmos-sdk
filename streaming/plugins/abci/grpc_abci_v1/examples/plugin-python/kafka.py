from concurrent import futures
import sys
import time
import grpc
import socket

import abci_listener_pb2
import abci_listener_pb2_grpc


from grpc_health.v1.health import HealthServicer
from grpc_health.v1 import health_pb2, health_pb2_grpc

from confluent_kafka import Producer

class ABCIListenerServiceServicer(abci_listener_pb2_grpc.ABCIListenerServiceServicer):
    """Implementation of ABCListener service."""

    producer = Producer({'bootstrap.servers': "localhost:9092",
                         'client.id': socket.gethostname()})

    def ListenBeginBlock(self, request, context):
        self.producer.produce("raw_begin_block_req", key=str(request.block_height), value=str(request.req))
        self.producer.produce("raw_begin_block_res", key=str(request.block_height), value=str(request.res))
        return abci_listener_pb2.Empty()

    def ListenEndBlock(self, request, context):
        self.producer.produce("raw_end_block_req", key=str(request.block_height), value=str(request.req))
        self.producer.produce("raw_end_block_res", key=str(request.block_height), value=str(request.res))
        return abci_listener_pb2.Empty()

    def ListenDeliverTx(self, request, context):
        key = "{}.{}".format(request.block_height, request.tx_idx)
        self. producer.produce("raw_deliver_tx_req", key=key, value=str(request.req))
        self.producer.produce("raw_deliver_tx_res", key=key, value=str(request.res))
        return abci_listener_pb2.Empty()

    def ListenStoreKVPair(self, request, context):
        key = "{}.{}".format(request.block_height, request.store_kv_pair_idx)
        self.producer.produce("raw_state_change", key=key, value=str(request.store_kv_pair))
        return abci_listener_pb2.Empty()

def serve():
    # We need to build a health service to work with go-plugin
    health = HealthServicer()
    health.set("plugin", health_pb2.HealthCheckResponse.ServingStatus.Value('SERVING'))

    # Start the server.
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    abci_listener_pb2_grpc.add_ABCIListenerServiceServicer_to_server(ABCIListenerServiceServicer(), server)
    health_pb2_grpc.add_HealthServicer_to_server(health, server)
    server.add_insecure_port('127.0.0.1:1234')
    server.start()

    # Output handshake information
    # https://github.com/hashicorp/go-plugin/blob/master/docs/guide-plugin-write-non-go.md#4-output-handshake-information
    print("1|1|tcp|127.0.0.1:1234|grpc")
    sys.stdout.flush()

    try:
        while True:
            time.sleep(60 * 60 * 24)
    except KeyboardInterrupt:
        server.stop(0)

if __name__ == '__main__':
    serve()