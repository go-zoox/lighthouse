(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([[88],{"fcC/":function(e,a,t){"use strict";t.r(a),t.d(a,"default",function(){return C});var n=t("p0pE"),r=t.n(n),c=t("qIgq"),i=t.n(c),l=t("q1tI"),o=t.n(l),d=t("ngQI"),s=t("R+Pm"),m=t("dhM6"),u=t("lfch"),g=t("/MfK"),p=t("WsrF"),b=(t("XlDN"),t("ZcRI")),f=t("ZTBq"),E=t.n(f),v=t("engV"),y=t("iW2X"),h=t("HD+4"),w=t("BbEZ"),j=t("fcXV"),k=t("PrJq"),I=t("8fPX"),N=Object(s["a"])({scriptUrl:"//at.alicdn.com/t/font_1507956_wwzrlcti59.js"});function A(e,a,t){var n=Array.from(e),r=n.splice(a,1),c=i()(r,1),l=c[0];return n.splice(t,0,l),n}function O(e,a,t,n,c){var i=Array.from(e),l=Array.from(a);console.log("createContainer: ",t,n);var o=i[t.index],d=Object(b["uuid"])(),s=r()({},o,{id:d,config:r()({},o.config,{value:r()({},o.config.value,{dataIndex:"".concat(o.type,"_").concat(d)})})});return l.splice(n.index,0,s),c(s.id),console.log(l),l}function S(e,a,t){var n=Array.from(e),c=r()({},a,{id:Object(b["uuid"])(),title:a.title+" Copy"});return n.splice(t+1,0,c),n}function x(e,a,t){var n=Array.from(e);return n.splice(t,1),n}function C(){var e=Object(l["useState"])(k["a"]),a=i()(e,2),t=a[0],n=(a[1],Object(l["useState"])([{id:"test_id",title:"\u5355\u884c\u6587\u5b57",type:"text",icon:"input",config:I["default"].text.config}])),c=i()(n,2),s=c[0],b=c[1],f=Object(l["useState"])({width:"100%",labelAlign:"right",labelWidth:4,labelHidden:!1,className:"",gap:16,disabled:!1,actionEnable:!1}),C=i()(f,2),q=C[0],_=C[1],R=Object(l["useState"])("test_id"),X=i()(R,2),B=X[0],D=X[1],H=s.find(e=>e.id===B),J=H,P=Object(l["useCallback"])((e,a)=>{var n=e.source,r=e.destination;if(r)switch(n.droppableId){case r.droppableId:b(A(s,n.index,r.index));break;case"images":b(O(t,s,n,r,e=>{D(e)}));break;case"containers":alert("Not Supported");break;default:alert("Not Supported");break}},[s]),W=(e,a)=>{if("component"===e){var t=B,n=s.map(e=>e.id!==t?e:r()({},e,a));b(n)}"form"===e&&_(a)},Z=e=>{return o.a.createElement(j["a"],Object.assign({},e.config,{type:e.type,global:q}))};return o.a.createElement("div",{className:"".concat(E.a["page"]," ").concat(E.a["images-container-box"]," ").concat(E.a.grid)},o.a.createElement(d["a"],{onDragEnd:P},o.a.createElement("div",{className:"".concat(E.a["grid-area"]," ").concat(E.a["grid-images"])},o.a.createElement(v["a"],{name:"images",dataSource:t,renderItem:e=>o.a.createElement("div",{style:{cursor:e.enable?"pointer":"not-allowed",backgroundColor:e.enable?"transparent":"rgba(0, 0, 0, .38)"}},o.a.createElement(N,{style:{marginRight:4},type:e.icon}),e.title)})),o.a.createElement("div",{className:"".concat(E.a["grid-area"]," ").concat(E.a["grid-containers"])},o.a.createElement("div",{className:E.a.actionBar},o.a.createElement(w["a"],null)),o.a.createElement("div",{className:E.a.contentArea},o.a.createElement(p["a"],{layout:"top"===q.labelAlign?"vertical":"horizontal"},o.a.createElement(y["a"],{name:"containers",dataSource:s,selectId:B,renderHandler:e=>o.a.createElement(m["a"],null),renderActions:(e,a)=>o.a.createElement(o.a.Fragment,null,o.a.createElement(u["a"],{style:{marginRight:6},onClick:()=>b(S(s,e,a))}),o.a.createElement(g["a"],{onClick:()=>b(x(s,e,a))})),renderItem:Z,onSelect:D})))),o.a.createElement("div",{className:"".concat(E.a["grid-area"]," ").concat(E.a["grid-attributes"])},o.a.createElement(h["a"],{component:J,form:q,onChange:W}))))}}}]);