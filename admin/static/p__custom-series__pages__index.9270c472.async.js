(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([[19],{"7nuH":function(e,a,n){"use strict";n.r(a);var t=n("p0pE"),i=n.n(t),o=n("q1tI"),r=n.n(o),g=n("/MKj"),s=n("RSRD"),c=n("W3hi");class l extends r.a.Component{render(){var e=this.props;return r.a.createElement("div",{style:{backgroundColor:"#fff"}},r.a.createElement(s["a"],Object.assign({loading:e.loading},e.attributes,{filters:e.attributes.filters.default,columns:e.attributes.tables.default,filterValues:e.constraints.filters,dataSource:e.datasets.table,onActionTrigger:e.onActionTrigger,onTableChange:e.onTableChange,onPaginationChange:e.onPaginationChange,onPaginationPageSizeChange:e.onPaginationPageSizeChange})))}}function u(e){var a=e.loading.models[c["e"]],n=e[c["e"]];return i()({loading:a},n)}function d(e){return{onActionTrigger:(a,n)=>e({type:"".concat(c["e"],"/dispatcher"),payload:{meta:a,data:n}}),onTableChange:(e,a,n)=>{},onPaginationChange:a=>e({type:"".concat(c["e"],"/service/query"),payload:{page:a}}),onPaginationPageSizeChange:(a,n)=>e({type:"".concat(c["e"],"/service/query/pageSize"),payload:{current:a,pageSize:n}})}}a["default"]=Object(g["c"])(u,d)(l)}}]);