(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([[42],{"+wfB":function(e,t,a){"use strict";a.r(t);var r=a("Y/ft"),n=a.n(r),s=a("d6i3"),c=a.n(s),i=a("p0pE"),u=a.n(i),p=a("gE9T"),o=a("V/h/"),d=a("LvDl"),l=a("sy1d"),y=a("VUo7");t["default"]={namespace:y["e"],state:{visible:{},attributes:{actions:y["a"],forms:y["d"],confirms:y["b"],filters:y["c"],tables:y["f"],pagination:{pageSizeOptions:["5","10","20"],current:1,pageSize:10,total:1}},datasets:{table:[],"create-dict-group::table":[]},constraints:{filters:{},orderBy:[]}},reducers:{save:(e,t)=>{var a=t.payload;return u()({},e,a)},"open/dialog":(e,t)=>{var a=t.payload;return Object(o["a"])(e,e=>{e.visible[a]=!0})},"close/dialog":(e,t)=>{var a=t.payload;return Object(o["a"])(e,e=>{e.visible[a]=!1})}},effects:{dispatcher(e,t){return c.a.mark(function a(){var r,n,s,i,u,p,o,d,l;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:if(r=e.payload,n=t.put,s=r||{},i=s.meta,u=void 0===i?{}:i,p=s.data,o=void 0===p?{}:p,d=u.effectType,l=void 0===d?"dialog":d,!u.effect){a.next=9;break}return a.next=7,n({type:u.effect,payload:o});case 7:a.next=11;break;case 9:return a.next=11,n({type:"".concat(l,"/").concat(u.key),payload:o});case 11:case"end":return a.stop()}},a)})()},"dialog/create"(e,t){return c.a.mark(function a(){var r,s,i,p,o,d,l;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return r=e.payload,s=t.select,i=t.put,a.next=4,s(e=>e[y["e"]]);case 4:return p=a.sent,o=p.attributes,d="create",r.values,r.initialValue,l=n()(r,["values","initialValue"]),a.next=10,i({type:"forms/open",payload:{attributes:u()({},o.forms[d],{target:{namespace:y["e"],effect:"service/create"}}),datasets:l}});case 10:case"end":return a.stop()}},a)})()},"dialog/update"(e,t){return c.a.mark(function a(){var r,n,s,i,p,o;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return r=e.payload.record,n=t.select,s=t.put,a.next=4,n(e=>e[y["e"]]);case 4:return i=a.sent,p=i.attributes,o="update",a.next=9,s({type:"forms/open",payload:{attributes:u()({},p.forms[o],{target:{namespace:y["e"],effect:"service/update"}}),datasets:r}});case 9:case"end":return a.stop()}},a)})()},"dialog/delete"(e,t){return c.a.mark(function a(){var r,n,s,i,p,o;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return e.type,r=e.payload.record,n=t.select,s=t.put,a.next=4,n(e=>e[y["e"]]);case 4:return i=a.sent,p=i.attributes,o="delete",a.next=9,s({type:"confirms/open",payload:{attributes:u()({},p.confirms[o],{target:{namespace:y["e"],effect:"service/delete"}}),datasets:r}});case 9:case"end":return a.stop()}},a)})()},"dialog/delete-dict-item"(e,t){return c.a.mark(function a(){var r,n,s,i,p,o;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return e.type,r=e.payload.record,n=t.select,s=t.put,a.next=4,n(e=>e[y["e"]]);case 4:return i=a.sent,p=i.attributes,o="delete",a.next=9,s({type:"confirms/open",payload:{attributes:u()({},p.confirms[o],{target:{namespace:y["e"],effect:"service/delete-dict-item"}}),datasets:r}});case 9:case"end":return a.stop()}},a)})()},"dialog/create-dict-group"(e,t){return c.a.mark(function a(){var r,n,s,i;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return e.type,r=e.payload.record,n=t.select,s=t.put,a.next=4,n(e=>e[y["e"]]);case 4:return i=a.sent,a.next=7,s({type:"service/query-dict-group",payload:r});case 7:return a.next=9,s({type:"save",payload:Object(o["a"])(i,e=>{e.visible["create-dict-group"]=!0,e.datasets["create-dict-group"]=r})});case 9:case"end":return a.stop()}},a)})()},"dialog/create-dict-item"(e,t){return c.a.mark(function a(){var r,n,s,i,p;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return r=e.payload,n=t.select,s=t.put,a.next=4,n(e=>e[y["e"]]);case 4:return i=a.sent,p=i.attributes,a.next=8,s({type:"forms/open",payload:{attributes:u()({},p.forms["create-dict-item"],{target:{namespace:y["e"],effect:"service/create-dict-item"}}),datasets:{group:r._id}}});case 8:case"end":return a.stop()}},a)})()},"service/query"(e,t){return c.a.mark(function a(){var r,n,s,i,d,f,m,v,x,g,w,b,h,k,q,O,E;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return r=e.payload,n=r.page,s=r.values,i=t.select,d=t.race,f=t.call,m=t.put,a.next=4,i(e=>e[y["e"]]);case 4:if(v=a.sent,!s){a.next=8;break}return a.next=8,m({type:"save",payload:{constraints:Object(o["a"])(v.constraints,e=>{e.filters=s})}});case 8:return a.next=10,i(e=>e[y["e"]]);case 10:return x=a.sent,g=x.attributes,w=x.datasets,b=x.constraints,h=g.pagination,k=n||h.current,a.next=18,d({data:f(l["a"],{method:"GET",path:"/api/resource/dicts",query:u()({page:k,pageSize:h.pageSize},b.filters)}),timeout:f(p["delay"],1e4)});case 18:if(q=a.sent,O=q.data,E=q.timeout,!E){a.next=25;break}return a.next=24,m({type:"messages/notice/error",payload:"\u670d\u52a1\u5668\u5f00\u5c0f\u5dee\u5566"});case 24:return a.abrupt("return",!1);case 25:return a.next=27,m({type:"save",payload:{attributes:Object(o["a"])(g,e=>{e.pagination.current=k,e.pagination.total=O.total}),datasets:Object(o["a"])(w,e=>{e.table=O.data})}});case 27:case"end":return a.stop()}},a)})()},"service/create"(e,t){return c.a.mark(function a(){var r,n,s,i,o,d,y;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return r=e.payload,n=t.race,s=t.call,i=t.put,o=u()({},r,{type:"group"}),a.next=5,n({data:s(l["a"],{method:"POST",path:"/api/resource/dicts",body:o}),timeout:s(p["delay"],1e4)});case 5:if(d=a.sent,d.data,y=d.timeout,!y){a.next=12;break}return a.next=11,i({type:"messages/tip/error",payload:"\u670d\u52a1\u5668\u5f00\u5c0f\u5dee\u5566"});case 11:return a.abrupt("return",!1);case 12:return a.next=14,i({type:"messages/tip/success",payload:"\u521b\u5efa\u6210\u529f"});case 14:return a.next=16,i({type:"service/query/first/page",payload:r});case 16:case"end":return a.stop()}},a)})()},"service/update"(e,t){return c.a.mark(function a(){var r,n,s,i,u,o,y;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return r=e.payload,n=t.race,s=t.call,i=t.put,u={id:r._id},a.next=5,n({data:s(l["a"],{method:"PUT",path:"/api/resource/dicts/:id",params:u,body:Object(d["pick"])(r,["name","description"])}),timeout:s(p["delay"],1e4)});case 5:if(o=a.sent,y=o.timeout,!y){a.next=11;break}return a.next=10,i({type:"messages/tip/error",payload:"\u670d\u52a1\u5668\u5f00\u5c0f\u5dee\u5566"});case 10:return a.abrupt("return",!1);case 11:return a.next=13,i({type:"messages/tip/success",payload:"\u66f4\u65b0\u6210\u529f"});case 13:return a.next=15,i({type:"service/query/first/page",payload:r});case 15:case"end":return a.stop()}},a)})()},"service/delete"(e,t){return c.a.mark(function a(){var r,n,s,i,u,o,d;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return r=e.payload,n=t.race,s=t.call,i=t.put,u={id:r._id},a.next=5,n({data:s(l["a"],{method:"DELETE",path:"/api/resource/dicts/:id",params:u}),timeout:s(p["delay"],1e4)});case 5:if(o=a.sent,d=o.timeout,!d){a.next=11;break}return a.next=10,i({type:"messages/tip/error",payload:"\u670d\u52a1\u5668\u5f00\u5c0f\u5dee\u5566"});case 10:return a.abrupt("return",!1);case 11:return a.next=13,i({type:"messages/tip/success",payload:"\u5220\u9664\u6210\u529f"});case 13:return a.next=15,i({type:"service/query/first/page",payload:r});case 15:case"end":return a.stop()}},a)})()},"service/query-dict-group"(e,t){return c.a.mark(function a(){var r,n,s,i,d,f,m,v,x;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return r=e.payload._id,n=t.select,s=t.race,i=t.call,d=t.put,a.next=4,n(e=>e[y["e"]]);case 4:return f=a.sent,a.next=7,s({data:i(l["a"],{method:"GET",path:"/api/resource/dicts/:id/items",params:{id:r}}),timeout:i(p["delay"],1e4)});case 7:if(m=a.sent,v=m.data,x=m.timeout,!x){a.next=14;break}return a.next=13,d({type:"messages/notice/error",payload:"\u670d\u52a1\u5668\u5f00\u5c0f\u5dee\u5566"});case 13:return a.abrupt("return",!1);case 14:return a.next=16,d({type:"save",payload:Object(o["a"])(f,e=>{e.datasets["create-dict-group::table"]=(v||[]).map(e=>u()({},e,{key:e._id}))})});case 16:case"end":return a.stop()}},a)})()},"service/create-dict-item"(e,t){return c.a.mark(function a(){var r,n,s,i,o,d,y;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return r=e.payload,n=t.race,s=t.call,i=t.put,o=u()({},r,{type:"item"}),a.next=5,n({data:s(l["a"],{method:"POST",path:"/api/resource/dicts",body:o}),timeout:s(p["delay"],1e4)});case 5:if(d=a.sent,d.data,y=d.timeout,!y){a.next=12;break}return a.next=11,i({type:"messages/tip/error",payload:"\u670d\u52a1\u5668\u5f00\u5c0f\u5dee\u5566"});case 11:return a.abrupt("return",!1);case 12:return a.next=14,i({type:"messages/tip/success",payload:"\u521b\u5efa\u6210\u529f"});case 14:return a.next=16,i({type:"service/query-dict-group",payload:{_id:r.group}});case 16:case"end":return a.stop()}},a)})()},"service/delete-dict-item"(e,t){return c.a.mark(function a(){var r,n,s,i,u,o,d;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return r=e.payload,n=t.race,s=t.call,i=t.put,u={id:r._id},a.next=5,n({data:s(l["a"],{method:"DELETE",path:"/api/resource/dicts/:id",params:u}),timeout:s(p["delay"],1e4)});case 5:if(o=a.sent,d=o.timeout,!d){a.next=11;break}return a.next=10,i({type:"messages/tip/error",payload:"\u670d\u52a1\u5668\u5f00\u5c0f\u5dee\u5566"});case 10:return a.abrupt("return",!1);case 11:return a.next=13,i({type:"messages/tip/success",payload:"\u5220\u9664\u6210\u529f"});case 13:return a.next=15,i({type:"service/query-dict-group",payload:{_id:r.group}});case 15:case"end":return a.stop()}},a)})()},"service/query/first/page"(e,t){return c.a.mark(function e(){var a;return c.a.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return a=t.put,e.next=3,a({type:"service/query",payload:{page:1}});case 3:case"end":return e.stop()}},e)})()},"service/query/refresh/page"(e,t){return c.a.mark(function e(){var a,r,n,s;return c.a.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return a=t.select,r=t.put,e.next=3,a(e=>e[y["e"]]);case 3:return n=e.sent,s=n.attributes,e.next=7,r({type:"service/query",payload:{page:s.pagination.current}});case 7:case"end":return e.stop()}},e)})()},"service/query/pageSize"(e,t){return c.a.mark(function a(){var r,n,s,i,u,p;return c.a.wrap(function(a){while(1)switch(a.prev=a.next){case 0:return r=e.payload,n=t.select,s=t.put,r.current,i=r.pageSize,a.next=5,n(e=>e[y["e"]]);case 5:return u=a.sent,p=u.attributes,a.next=9,s({type:"save",payload:{attributes:Object(o["a"])(p,e=>{e.pagination.current=1,e.pagination.pageSize=i})}});case 9:return a.next=11,s({type:"service/query",payload:{page:1}});case 11:case"end":return a.stop()}},a)})()},initialize(e,t){return c.a.mark(function e(){var a,r;return c.a.wrap(function(e){while(1)switch(e.prev=e.next){case 0:return a=t.put,r=t.take,e.next=3,a({type:"service/query/first/page"});case 3:return e.next=5,r("service/query");case 5:case"end":return e.stop()}},e)})()}},subscriptions:{}}}}]);