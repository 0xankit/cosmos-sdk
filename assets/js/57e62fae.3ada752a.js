"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[3458],{3905:(e,n,t)=>{t.d(n,{Zo:()=>s,kt:()=>m});var r=t(7294);function a(e,n,t){return n in e?Object.defineProperty(e,n,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[n]=t,e}function o(e,n){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);n&&(r=r.filter((function(n){return Object.getOwnPropertyDescriptor(e,n).enumerable}))),t.push.apply(t,r)}return t}function i(e){for(var n=1;n<arguments.length;n++){var t=null!=arguments[n]?arguments[n]:{};n%2?o(Object(t),!0).forEach((function(n){a(e,n,t[n])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(t)):o(Object(t)).forEach((function(n){Object.defineProperty(e,n,Object.getOwnPropertyDescriptor(t,n))}))}return e}function p(e,n){if(null==e)return{};var t,r,a=function(e,n){if(null==e)return{};var t,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)t=o[r],n.indexOf(t)>=0||(a[t]=e[t]);return a}(e,n);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)t=o[r],n.indexOf(t)>=0||Object.prototype.propertyIsEnumerable.call(e,t)&&(a[t]=e[t])}return a}var c=r.createContext({}),l=function(e){var n=r.useContext(c),t=n;return e&&(t="function"==typeof e?e(n):i(i({},n),e)),t},s=function(e){var n=l(e.components);return r.createElement(c.Provider,{value:n},e.children)},d={inlineCode:"code",wrapper:function(e){var n=e.children;return r.createElement(r.Fragment,{},n)}},u=r.forwardRef((function(e,n){var t=e.components,a=e.mdxType,o=e.originalType,c=e.parentName,s=p(e,["components","mdxType","originalType","parentName"]),u=l(t),m=a,k=u["".concat(c,".").concat(m)]||u[m]||d[m]||o;return t?r.createElement(k,i(i({ref:n},s),{},{components:t})):r.createElement(k,i({ref:n},s))}));function m(e,n){var t=arguments,a=n&&n.mdxType;if("string"==typeof e||a){var o=t.length,i=new Array(o);i[0]=u;var p={};for(var c in n)hasOwnProperty.call(n,c)&&(p[c]=n[c]);p.originalType=e,p.mdxType="string"==typeof e?e:a,i[1]=p;for(var l=2;l<o;l++)i[l]=t[l];return r.createElement.apply(null,i)}return r.createElement.apply(null,t)}u.displayName="MDXCreateElement"},9045:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>c,contentTitle:()=>i,default:()=>d,frontMatter:()=>o,metadata:()=>p,toc:()=>l});var r=t(7462),a=(t(7294),t(3905));const o={sidebar_position:1},i="Depinject",p={unversionedId:"tooling/depinject",id:"version-v0.47/tooling/depinject",title:"Depinject",description:"DISCLAIMER: This is a beta package. The SDK team is actively working on this feature and we are looking for feedback from the community. Please try it out and let us know what you think.",source:"@site/versioned_docs/version-v0.47/tooling/02-depinject.md",sourceDirName:"tooling",slug:"/tooling/depinject",permalink:"/v0.47/tooling/depinject",draft:!1,tags:[],version:"v0.47",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Cosmovisor",permalink:"/v0.47/tooling/cosmovisor"},next:{title:"Architecture Decision Records (ADR)",permalink:"/v0.47/architecture/"}},c={},l=[{value:"Overview",id:"overview",level:2},{value:"Usage",id:"usage",level:2},{value:"Interface type resolution",id:"interface-type-resolution",level:3},{value:"<code>BindInterface</code> API",id:"bindinterface-api",level:4},{value:"Full example in real app",id:"full-example-in-real-app",level:3},{value:"Debugging",id:"debugging",level:2}],s={toc:l};function d(e){let{components:n,...t}=e;return(0,a.kt)("wrapper",(0,r.Z)({},s,t,{components:n,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"depinject"},"Depinject"),(0,a.kt)("blockquote",null,(0,a.kt)("p",{parentName:"blockquote"},(0,a.kt)("strong",{parentName:"p"},"DISCLAIMER"),": This is a ",(0,a.kt)("strong",{parentName:"p"},"beta")," package. The SDK team is actively working on this feature and we are looking for feedback from the community. Please try it out and let us know what you think.")),(0,a.kt)("h2",{id:"overview"},"Overview"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"depinject")," is a dependency injection framework for the Cosmos SDK. This module together with ",(0,a.kt)("inlineCode",{parentName:"p"},"core/appconfig")," are meant to simplify the definition of a blockchain by replacing most of ",(0,a.kt)("inlineCode",{parentName:"p"},"app.go"),"'s boilerplate code with a configuration file (Go, YAML or JSON)."),(0,a.kt)("ul",null,(0,a.kt)("li",{parentName:"ul"},(0,a.kt)("a",{parentName:"li",href:"https://pkg.go.dev/cosmossdk.io/depinject"},"Go Doc"))),(0,a.kt)("h2",{id:"usage"},"Usage"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"depinject")," includes an expressive and composable ",(0,a.kt)("a",{parentName:"p",href:"https://pkg.go.dev/github.com/cosmos/cosmos-sdk/depinject#Config"},"Configuration API"),".\nA core configuration function is ",(0,a.kt)("inlineCode",{parentName:"p"},"Provide"),". The example below demonstrates the registration of free ",(0,a.kt)("strong",{parentName:"p"},"provider functions")," via the ",(0,a.kt)("inlineCode",{parentName:"p"},"Provide")," API."),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'package main\n\nimport (\n    "fmt"\n\n    "cosmossdk.io/depinject"\n)\n\ntype AnotherInt int\n\nfunc main() {\n    var (\n      x int\n      y AnotherInt\n    )\n\n    fmt.Printf("Before (%v, %v)\\n", x, y)\n    depinject.Inject(\n        depinject.Provide(\n            func() int { return 1 },\n            func() AnotherInt { return AnotherInt(2) },\n        ),\n        &x,\n        &y,\n    )\n    fmt.Printf("After (%v, %v)\\n", x, y)\n}\n')),(0,a.kt)("p",null,"Provider functions form the basis of the dependency tree, they are introspected then their inputs identified as dependencies and outputs as dependants, either for another provider function or state stored outside the DI container, as is the case of ",(0,a.kt)("inlineCode",{parentName:"p"},"&x")," and ",(0,a.kt)("inlineCode",{parentName:"p"},"&y")," above."),(0,a.kt)("h3",{id:"interface-type-resolution"},"Interface type resolution"),(0,a.kt)("p",null,(0,a.kt)("inlineCode",{parentName:"p"},"depinject")," supports interface types as inputs to provider functions.  In the SDK's case this pattern is used to decouple\n",(0,a.kt)("inlineCode",{parentName:"p"},"Keeper")," dependencies between modules.  For example ",(0,a.kt)("inlineCode",{parentName:"p"},"x/bank")," expects an ",(0,a.kt)("a",{parentName:"p",href:"https://pkg.go.dev/github.com/cosmos/cosmos-sdk/x/bank/types#AccountKeeper"},"AccountKeeper")," interface as ",(0,a.kt)("a",{parentName:"p",href:"https://github.com/cosmos/cosmos-sdk/blob/de343d458aa68c19630177807d6f0e2e6deaf7a9/x/bank/module.go#L224"},"input to provideModule"),"."),(0,a.kt)("p",null,"Concretely ",(0,a.kt)("inlineCode",{parentName:"p"},"SimApp")," uses the implementation in ",(0,a.kt)("inlineCode",{parentName:"p"},"x/auth"),", but this design allows for this loose coupling to change."),(0,a.kt)("p",null,"Given the following types:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"package duck\n\ntype Duck interface {\n    quack()\n}\n\ntype AlsoDuck interface {\n    quack()\n}\n\ntype Mallard struct{}\ntype Canvasback struct{}\n\nfunc (duck Mallard) quack()    {}\nfunc (duck Canvasback) quack() {}\n\ntype Pond struct {\n    Duck AlsoDuck\n}\n")),(0,a.kt)("p",null,"This usage"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"var pond Pond\n\ndepinject.Inject(\n  depinject.Provide(\n    func() Mallard { return Mallard{} },\n    func(duck Duck) Pond {\n      return Pond{Duck: duck}\n    }),\n   &pond)\n")),(0,a.kt)("p",null,"results in an ",(0,a.kt)("em",{parentName:"p"},"implicit")," binding of ",(0,a.kt)("inlineCode",{parentName:"p"},"Duck")," to ",(0,a.kt)("inlineCode",{parentName:"p"},"Mallard"),".  This works because there is only one implementation of ",(0,a.kt)("inlineCode",{parentName:"p"},"Duck")," in the container.",(0,a.kt)("br",{parentName:"p"}),"\n","However, adding a second provider of ",(0,a.kt)("inlineCode",{parentName:"p"},"Duck")," will result in an error:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"var pond Pond\n\ndepinject.Inject(\n  depinject.Provide(\n    func() Mallard { return Mallard{} },\n    func() Canvasback { return Canvasback{} },\n    func(duck Duck) Pond {\n      return Pond{Duck: duck}\n    }),\n   &pond)\n")),(0,a.kt)("p",null,"A specific binding preference for ",(0,a.kt)("inlineCode",{parentName:"p"},"Duck")," is required."),(0,a.kt)("h4",{id:"bindinterface-api"},(0,a.kt)("inlineCode",{parentName:"h4"},"BindInterface")," API"),(0,a.kt)("p",null,"In the above situation registering a binding for a given interface binding may look like"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'depinject.Inject(\n  depinject.Configs(\n    depinject.BindInterface(\n      "duck.Duck",\n      "duck.Mallard"),\n     depinject.Provide(\n       func() Mallard { return Mallard{} },\n       func() Canvasback { return Canvasback{} },\n       func(duck Duck) APond {\n         return Pond{Duck: duck}\n      })),\n   &pond)\n')),(0,a.kt)("p",null,"Now ",(0,a.kt)("inlineCode",{parentName:"p"},"depinject")," has enough information to provide ",(0,a.kt)("inlineCode",{parentName:"p"},"Mallard")," as an input to ",(0,a.kt)("inlineCode",{parentName:"p"},"APond"),". "),(0,a.kt)("h3",{id:"full-example-in-real-app"},"Full example in real app"),(0,a.kt)("admonition",{type:"warning"},(0,a.kt)("p",{parentName:"admonition"},"When using ",(0,a.kt)("inlineCode",{parentName:"p"},"depinject.Inject"),", the injected types must be pointers.")),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go",metastring:"reference",reference:!0},"https://github.com/cosmos/cosmos-sdk/blob/v0.47.0-alpha1/simapp/app.go#L241-L266\n")),(0,a.kt)("h2",{id:"debugging"},"Debugging"),(0,a.kt)("p",null,"Issues with resolving dependencies in the container can be done with logs and ",(0,a.kt)("a",{parentName:"p",href:"https://graphviz.org"},"Graphviz")," renderings of the container tree.\nBy default, whenever there is an error, logs will be printed to stderr and a rendering of the dependency graph in Graphviz DOT format will be saved to ",(0,a.kt)("inlineCode",{parentName:"p"},"debug_container.dot"),"."),(0,a.kt)("p",null,"Here is an example Graphviz rendering of a successful build of a dependency graph:\n",(0,a.kt)("img",{parentName:"p",src:"https://raw.githubusercontent.com/cosmos/cosmos-sdk/ff39d243d421442b400befcd959ec3ccd2525154/depinject/testdata/example.svg",alt:"Graphviz Example"})),(0,a.kt)("p",null,"Rectangles represent functions, ovals represent types, rounded rectangles represent modules and the single hexagon\nrepresents the function which called ",(0,a.kt)("inlineCode",{parentName:"p"},"Build"),". Black-colored shapes mark functions and types that were called/resolved\nwithout an error. Gray-colored nodes mark functions and types that could have been called/resolved in the container but\nwere left unused."),(0,a.kt)("p",null,"Here is an example Graphviz rendering of a dependency graph build which failed:\n",(0,a.kt)("img",{parentName:"p",src:"https://raw.githubusercontent.com/cosmos/cosmos-sdk/ff39d243d421442b400befcd959ec3ccd2525154/depinject/testdata/example_error.svg",alt:"Graphviz Error Example"})),(0,a.kt)("p",null,"Graphviz DOT files can be converted into SVG's for viewing in a web browser using the ",(0,a.kt)("inlineCode",{parentName:"p"},"dot")," command-line tool, ex:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-txt"},"dot -Tsvg debug_container.dot > debug_container.svg\n")),(0,a.kt)("p",null,"Many other tools including some IDEs support working with DOT files."))}d.isMDXComponent=!0}}]);