"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[3496],{3905:(e,t,o)=>{o.d(t,{Zo:()=>p,kt:()=>c});var n=o(7294);function a(e,t,o){return t in e?Object.defineProperty(e,t,{value:o,enumerable:!0,configurable:!0,writable:!0}):e[t]=o,e}function i(e,t){var o=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),o.push.apply(o,n)}return o}function s(e){for(var t=1;t<arguments.length;t++){var o=null!=arguments[t]?arguments[t]:{};t%2?i(Object(o),!0).forEach((function(t){a(e,t,o[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(o)):i(Object(o)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(o,t))}))}return e}function r(e,t){if(null==e)return{};var o,n,a=function(e,t){if(null==e)return{};var o,n,a={},i=Object.keys(e);for(n=0;n<i.length;n++)o=i[n],t.indexOf(o)>=0||(a[o]=e[o]);return a}(e,t);if(Object.getOwnPropertySymbols){var i=Object.getOwnPropertySymbols(e);for(n=0;n<i.length;n++)o=i[n],t.indexOf(o)>=0||Object.prototype.propertyIsEnumerable.call(e,o)&&(a[o]=e[o])}return a}var l=n.createContext({}),d=function(e){var t=n.useContext(l),o=t;return e&&(o="function"==typeof e?e(t):s(s({},t),e)),o},p=function(e){var t=d(e.components);return n.createElement(l.Provider,{value:t},e.children)},m={inlineCode:"code",wrapper:function(e){var t=e.children;return n.createElement(n.Fragment,{},t)}},u=n.forwardRef((function(e,t){var o=e.components,a=e.mdxType,i=e.originalType,l=e.parentName,p=r(e,["components","mdxType","originalType","parentName"]),u=d(o),c=a,h=u["".concat(l,".").concat(c)]||u[c]||m[c]||i;return o?n.createElement(h,s(s({ref:t},p),{},{components:o})):n.createElement(h,s({ref:t},p))}));function c(e,t){var o=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var i=o.length,s=new Array(i);s[0]=u;var r={};for(var l in t)hasOwnProperty.call(t,l)&&(r[l]=t[l]);r.originalType=e,r.mdxType="string"==typeof e?e:a,s[1]=r;for(var d=2;d<i;d++)s[d]=o[d];return n.createElement.apply(null,s)}return n.createElement.apply(null,o)}u.displayName="MDXCreateElement"},6346:(e,t,o)=>{o.r(t),o.d(t,{assets:()=>l,contentTitle:()=>s,default:()=>m,frontMatter:()=>i,metadata:()=>r,toc:()=>d});var n=o(7462),a=(o(7294),o(3905));const i={sidebar_position:1},s="Introduction to Cosmos SDK Modules",r={unversionedId:"building-modules/intro",id:"version-v0.47/building-modules/intro",title:"Introduction to Cosmos SDK Modules",description:"Modules define most of the logic of Cosmos SDK applications. Developers compose modules together using the Cosmos SDK to build their custom application-specific blockchains. This document outlines the basic concepts behind SDK modules and how to approach module management.",source:"@site/versioned_docs/version-v0.47/building-modules/01-intro.md",sourceDirName:"building-modules",slug:"/building-modules/intro",permalink:"/v0.47/building-modules/intro",draft:!1,tags:[],version:"v0.47",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"In-Place Store Migrations",permalink:"/v0.47/core/upgrade"},next:{title:"Module Manager",permalink:"/v0.47/building-modules/module-manager"}},l={},d=[{value:"Role of Modules in a Cosmos SDK Application",id:"role-of-modules-in-a-cosmos-sdk-application",level:2},{value:"How to Approach Building Modules as a Developer",id:"how-to-approach-building-modules-as-a-developer",level:2},{value:"Main Components of Cosmos SDK Modules",id:"main-components-of-cosmos-sdk-modules",level:2}],p={toc:d};function m(e){let{components:t,...o}=e;return(0,a.kt)("wrapper",(0,n.Z)({},p,o,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"introduction-to-cosmos-sdk-modules"},"Introduction to Cosmos SDK Modules"),(0,a.kt)("admonition",{title:"Synopsis",type:"note"},(0,a.kt)("p",{parentName:"admonition"},"Modules define most of the logic of Cosmos SDK applications. Developers compose modules together using the Cosmos SDK to build their custom application-specific blockchains. This document outlines the basic concepts behind SDK modules and how to approach module management.")),(0,a.kt)("admonition",{type:"note"},(0,a.kt)("h3",{parentName:"admonition",id:"pre-requisite-readings"},"Pre-requisite Readings"),(0,a.kt)("ul",{parentName:"admonition"},(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/v0.47/basics/app-anatomy"},"Anatomy of a Cosmos SDK application")),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"/v0.47/basics/tx-lifecycle"},"Lifecycle of a Cosmos SDK transaction")))),(0,a.kt)("h2",{id:"role-of-modules-in-a-cosmos-sdk-application"},"Role of Modules in a Cosmos SDK Application"),(0,a.kt)("p",null,"The Cosmos SDK can be thought of as the Ruby-on-Rails of blockchain development. It comes with a core that provides the basic functionalities every blockchain application needs, like a ",(0,a.kt)("a",{parentName:"p",href:"/v0.47/core/baseapp"},"boilerplate implementation of the ABCI")," to communicate with the underlying consensus engine, a ",(0,a.kt)("a",{parentName:"p",href:"/v0.47/core/store#multistore"},(0,a.kt)("inlineCode",{parentName:"a"},"multistore"))," to persist state, a ",(0,a.kt)("a",{parentName:"p",href:"/v0.47/core/node"},"server")," to form a full-node and ",(0,a.kt)("a",{parentName:"p",href:"/v0.47/building-modules/module-interfaces"},"interfaces")," to handle queries."),(0,a.kt)("p",null,"On top of this core, the Cosmos SDK enables developers to build modules that implement the business logic of their application. In other words, SDK modules implement the bulk of the logic of applications, while the core does the wiring and enables modules to be composed together. The end goal is to build a robust ecosystem of open-source Cosmos SDK modules, making it increasingly easier to build complex blockchain applications."),(0,a.kt)("p",null,"Cosmos SDK modules can be seen as little state-machines within the state-machine. They generally define a subset of the state using one or more ",(0,a.kt)("inlineCode",{parentName:"p"},"KVStore"),"s in the ",(0,a.kt)("a",{parentName:"p",href:"/v0.47/core/store"},"main multistore"),", as well as a subset of ",(0,a.kt)("a",{parentName:"p",href:"/v0.47/building-modules/messages-and-queries#messages"},"message types"),". These messages are routed by one of the main components of Cosmos SDK core, ",(0,a.kt)("a",{parentName:"p",href:"/v0.47/core/baseapp"},(0,a.kt)("inlineCode",{parentName:"a"},"BaseApp")),", to a module Protobuf ",(0,a.kt)("a",{parentName:"p",href:"/v0.47/building-modules/msg-services"},(0,a.kt)("inlineCode",{parentName:"a"},"Msg")," service")," that defines them."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-text"},"                                      +\n                                      |\n                                      |  Transaction relayed from the full-node's consensus engine\n                                      |  to the node's application via DeliverTx\n                                      |\n                                      |\n                                      |\n                +---------------------v--------------------------+\n                |                 APPLICATION                    |\n                |                                                |\n                |     Using baseapp's methods: Decode the Tx,    |\n                |     extract and route the message(s)           |\n                |                                                |\n                +---------------------+--------------------------+\n                                      |\n                                      |\n                                      |\n                                      +---------------------------+\n                                                                  |\n                                                                  |\n                                                                  |\n                                                                  |  Message routed to the correct\n                                                                  |  module to be processed\n                                                                  |\n                                                                  |\n+----------------+  +---------------+  +----------------+  +------v----------+\n|                |  |               |  |                |  |                 |\n|  AUTH MODULE   |  |  BANK MODULE  |  | STAKING MODULE |  |   GOV MODULE    |\n|                |  |               |  |                |  |                 |\n|                |  |               |  |                |  | Handles message,|\n|                |  |               |  |                |  | Updates state   |\n|                |  |               |  |                |  |                 |\n+----------------+  +---------------+  +----------------+  +------+----------+\n                                                                  |\n                                                                  |\n                                                                  |\n                                                                  |\n                                       +--------------------------+\n                                       |\n                                       | Return result to the underlying consensus engine (e.g. Tendermint)\n                                       | (0=Ok, 1=Err)\n                                       v\n")),(0,a.kt)("p",null,"As a result of this architecture, building a Cosmos SDK application usually revolves around writing modules to implement the specialized logic of the application and composing them with existing modules to complete the application. Developers will generally work on modules that implement logic needed for their specific use case that do not exist yet, and will use existing modules for more generic functionalities like staking, accounts, or token management."),(0,a.kt)("h2",{id:"how-to-approach-building-modules-as-a-developer"},"How to Approach Building Modules as a Developer"),(0,a.kt)("p",null,"While there are no definitive guidelines for writing modules, here are some important design principles developers should keep in mind when building them:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("strong",{parentName:"li"},"Composability"),": Cosmos SDK applications are almost always composed of multiple modules. This means developers need to carefully consider the integration of their module not only with the core of the Cosmos SDK, but also with other modules. The former is achieved by following standard design patterns outlined ",(0,a.kt)("a",{parentName:"li",href:"#main-components-of-sdk-modules"},"here"),", while the latter is achieved by properly exposing the store(s) of the module via the ",(0,a.kt)("a",{parentName:"li",href:"/v0.47/building-modules/keeper"},(0,a.kt)("inlineCode",{parentName:"a"},"keeper")),"."),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("strong",{parentName:"li"},"Specialization"),": A direct consequence of the ",(0,a.kt)("strong",{parentName:"li"},"composability")," feature is that modules should be ",(0,a.kt)("strong",{parentName:"li"},"specialized"),". Developers should carefully establish the scope of their module and not batch multiple functionalities into the same module. This separation of concerns enables modules to be re-used in other projects and improves the upgradability of the application. ",(0,a.kt)("strong",{parentName:"li"},"Specialization")," also plays an important role in the ",(0,a.kt)("a",{parentName:"li",href:"/v0.47/core/ocap"},"object-capabilities model")," of the Cosmos SDK."),(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("strong",{parentName:"li"},"Capabilities"),": Most modules need to read and/or write to the store(s) of other modules. However, in an open-source environment, it is possible for some modules to be malicious. That is why module developers need to carefully think not only about how their module interacts with other modules, but also about how to give access to the module's store(s). The Cosmos SDK takes a capabilities-oriented approach to inter-module security. This means that each store defined by a module is accessed by a ",(0,a.kt)("inlineCode",{parentName:"li"},"key"),", which is held by the module's ",(0,a.kt)("a",{parentName:"li",href:"/v0.47/building-modules/keeper"},(0,a.kt)("inlineCode",{parentName:"a"},"keeper")),". This ",(0,a.kt)("inlineCode",{parentName:"li"},"keeper")," defines how to access the store(s) and under what conditions. Access to the module's store(s) is done by passing a reference to the module's ",(0,a.kt)("inlineCode",{parentName:"li"},"keeper"),".")),(0,a.kt)("h2",{id:"main-components-of-cosmos-sdk-modules"},"Main Components of Cosmos SDK Modules"),(0,a.kt)("p",null,"Modules are by convention defined in the ",(0,a.kt)("inlineCode",{parentName:"p"},"./x/")," subfolder (e.g. the ",(0,a.kt)("inlineCode",{parentName:"p"},"bank")," module will be defined in the ",(0,a.kt)("inlineCode",{parentName:"p"},"./x/bank")," folder). They generally share the same core components:"),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},"A  ",(0,a.kt)("a",{parentName:"li",href:"/v0.47/building-modules/keeper"},(0,a.kt)("inlineCode",{parentName:"a"},"keeper")),", used to access the module's store(s) and update the state."),(0,a.kt)("li",{parentName:"ul"},"A ",(0,a.kt)("a",{parentName:"li",href:"/v0.47/building-modules/messages-and-queries#messages"},(0,a.kt)("inlineCode",{parentName:"a"},"Msg")," service"),", used to process messages when they are routed to the module by ",(0,a.kt)("a",{parentName:"li",href:"/v0.47/core/baseapp#message-routing"},(0,a.kt)("inlineCode",{parentName:"a"},"BaseApp"))," and trigger state-transitions."),(0,a.kt)("li",{parentName:"ul"},"A ",(0,a.kt)("a",{parentName:"li",href:"/v0.47/building-modules/query-services"},"query service"),", used to process user queries when they are routed to the module by ",(0,a.kt)("a",{parentName:"li",href:"/v0.47/core/baseapp#query-routing"},(0,a.kt)("inlineCode",{parentName:"a"},"BaseApp")),"."),(0,a.kt)("li",{parentName:"ul"},"Interfaces, for end users to query the subset of the state defined by the module and create ",(0,a.kt)("inlineCode",{parentName:"li"},"message"),"s of the custom types defined in the module.")),(0,a.kt)("p",null,"In addition to these components, modules implement the ",(0,a.kt)("inlineCode",{parentName:"p"},"AppModule")," interface in order to be managed by the ",(0,a.kt)("a",{parentName:"p",href:"/v0.47/building-modules/module-manager"},(0,a.kt)("inlineCode",{parentName:"a"},"module manager")),"."),(0,a.kt)("p",null,"Please refer to the ",(0,a.kt)("a",{parentName:"p",href:"/v0.47/building-modules/structure"},"structure document")," to learn about the recommended structure of a module's directory."))}m.isMDXComponent=!0}}]);