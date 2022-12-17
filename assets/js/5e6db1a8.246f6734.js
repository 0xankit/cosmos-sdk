"use strict";(self.webpackChunkcosmos_sdk_docs=self.webpackChunkcosmos_sdk_docs||[]).push([[278],{3905:(e,t,n)=>{n.d(t,{Zo:()=>m,kt:()=>c});var r=n(7294);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function p(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var d=r.createContext({}),l=function(e){var t=r.useContext(d),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},m=function(e){var t=l(e.components);return r.createElement(d.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},s=r.forwardRef((function(e,t){var n=e.components,a=e.mdxType,o=e.originalType,d=e.parentName,m=p(e,["components","mdxType","originalType","parentName"]),s=l(n),c=a,g=s["".concat(d,".").concat(c)]||s[c]||u[c]||o;return n?r.createElement(g,i(i({ref:t},m),{},{components:n})):r.createElement(g,i({ref:t},m))}));function c(e,t){var n=arguments,a=t&&t.mdxType;if("string"==typeof e||a){var o=n.length,i=new Array(o);i[0]=s;var p={};for(var d in t)hasOwnProperty.call(t,d)&&(p[d]=t[d]);p.originalType=e,p.mdxType="string"==typeof e?e:a,i[1]=p;for(var l=2;l<o;l++)i[l]=n[l];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}s.displayName="MDXCreateElement"},5545:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>d,contentTitle:()=>i,default:()=>u,frontMatter:()=>o,metadata:()=>p,toc:()=>l});var r=n(7462),a=(n(7294),n(3905));const o={sidebar_position:1},i="Pre-Upgrade Handling",p={unversionedId:"migrations/pre-upgrade",id:"version-v0.47/migrations/pre-upgrade",title:"Pre-Upgrade Handling",description:"Cosmovisor supports custom pre-upgrade handling. Use pre-upgrade handling when you need to implement application config changes that are required in the newer version before you perform the upgrade.",source:"@site/versioned_docs/version-v0.47/migrations/01-pre-upgrade.md",sourceDirName:"migrations",slug:"/migrations/pre-upgrade",permalink:"/v0.47/migrations/pre-upgrade",draft:!1,tags:[],version:"v0.47",sidebarPosition:1,frontMatter:{sidebar_position:1},sidebar:"tutorialSidebar",previous:{title:"Introduction to Migrations",permalink:"/v0.47/migrations/intro"},next:{title:"List of Modules",permalink:"/v0.47/modules"}},d={},l=[{value:"Sample",id:"sample",level:2}],m={toc:l};function u(e){let{components:t,...n}=e;return(0,a.kt)("wrapper",(0,r.Z)({},m,n,{components:t,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"pre-upgrade-handling"},"Pre-Upgrade Handling"),(0,a.kt)("p",null,"Cosmovisor supports custom pre-upgrade handling. Use pre-upgrade handling when you need to implement application config changes that are required in the newer version before you perform the upgrade."),(0,a.kt)("p",null,"Using Cosmovisor pre-upgrade handling is optional. If pre-upgrade handling is not implemented, the upgrade continues."),(0,a.kt)("p",null,"For example, make the required new-version changes to ",(0,a.kt)("inlineCode",{parentName:"p"},"app.toml")," settings during the pre-upgrade handling. The pre-upgrade handling process means that the file does not have to be manually updated after the upgrade."),(0,a.kt)("p",null,"Before the application binary is upgraded, Cosmovisor calls a ",(0,a.kt)("inlineCode",{parentName:"p"},"pre-upgrade")," command that can  be implemented by the application."),(0,a.kt)("p",null,"The ",(0,a.kt)("inlineCode",{parentName:"p"},"pre-upgrade")," command does not take in any command-line arguments and is expected to terminate with the following exit codes:"),(0,a.kt)("table",null,(0,a.kt)("thead",{parentName:"table"},(0,a.kt)("tr",{parentName:"thead"},(0,a.kt)("th",{parentName:"tr",align:null},"Exit status code"),(0,a.kt)("th",{parentName:"tr",align:null},"How it is handled in Cosmosvisor"))),(0,a.kt)("tbody",{parentName:"table"},(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"0")),(0,a.kt)("td",{parentName:"tr",align:null},"Assumes ",(0,a.kt)("inlineCode",{parentName:"td"},"pre-upgrade")," command executed successfully and continues the upgrade.")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"1")),(0,a.kt)("td",{parentName:"tr",align:null},"Default exit code when ",(0,a.kt)("inlineCode",{parentName:"td"},"pre-upgrade")," command has not been implemented.")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"30")),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"pre-upgrade")," command was executed but failed. This fails the entire upgrade.")),(0,a.kt)("tr",{parentName:"tbody"},(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"31")),(0,a.kt)("td",{parentName:"tr",align:null},(0,a.kt)("inlineCode",{parentName:"td"},"pre-upgrade")," command was executed but failed. But the command is retried until exit code ",(0,a.kt)("inlineCode",{parentName:"td"},"1")," or ",(0,a.kt)("inlineCode",{parentName:"td"},"30")," are returned.")))),(0,a.kt)("h2",{id:"sample"},"Sample"),(0,a.kt)("p",null,"Here is a sample structure of the ",(0,a.kt)("inlineCode",{parentName:"p"},"pre-upgrade")," command:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},'func preUpgradeCommand() *cobra.Command {\n    cmd := &cobra.Command{\n        Use:   "pre-upgrade",\n        Short: "Pre-upgrade command",\n        Long: "Pre-upgrade command to implement custom pre-upgrade handling",\n        Run: func(cmd *cobra.Command, args []string) {\n\n            err := HandlePreUpgrade()\n\n            if err != nil {\n                os.Exit(30)\n            }\n\n            os.Exit(0)\n\n        },\n    }\n\n    return cmd\n}\n')),(0,a.kt)("p",null,"Ensure that the pre-upgrade command has been registered in the application:"),(0,a.kt)("pre",null,(0,a.kt)("code",{parentName:"pre",className:"language-go"},"rootCmd.AddCommand(\n        // ..\n        preUpgradeCommand(),\n        // ..\n    )\n")))}u.isMDXComponent=!0}}]);