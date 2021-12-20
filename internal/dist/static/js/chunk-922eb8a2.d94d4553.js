(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-922eb8a2"],{"0e54":function(t,e,a){},"1fe5":function(t,e,a){"use strict";a("0e54")},"93dc":function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"app-container"},[a("div",{staticClass:"filter-container fr"},[a("el-button",{staticClass:"filter-item",attrs:{icon:"el-icon-refresh",size:"small"},on:{click:t.onRefresh}}),a("router-link",{attrs:{to:{name:"createTask"},tag:"span"}},[a("el-button",{staticClass:"filter-item",attrs:{type:"primary",icon:"el-icon-edit",size:"small"}},[t._v("创建任务")])],1)],1),a("el-table",{directives:[{name:"loading",rawName:"v-loading",value:t.listLoading,expression:"listLoading"}],attrs:{data:t.list,"element-loading-text":"Loading",border:"",fit:"","highlight-current-row":""}},[a("el-table-column",{attrs:{align:"center",label:"编号",width:"60"},scopedSlots:t._u([{key:"default",fn:function(e){var a=e.row;return[t._v(" "+t._s(a.id)+" ")]}}])}),a("el-table-column",{attrs:{align:"center",label:"名称",width:"170"},scopedSlots:t._u([{key:"default",fn:function(e){var a=e.row;return[t._v(" "+t._s(a.task_name)+" ")]}}])}),a("el-table-column",{attrs:{align:"center","class-name":"status-col",label:"状态",width:"100"},scopedSlots:t._u([{key:"default",fn:function(e){var n=e.row;return[a("el-tag",[t._v(t._s(t._f("statusFilter")(n.status)))])]}}])}),a("el-table-column",{attrs:{align:"center",label:"运行次数",width:"100"},scopedSlots:t._u([{key:"default",fn:function(e){var a=e.row;return[t._v(" "+t._s(a.counts)+" ")]}}])}),a("el-table-column",{attrs:{align:"center",label:"定时任务",width:"80"},scopedSlots:t._u([{key:"default",fn:function(e){var a=e.row;return[t._v(" "+t._s(a.isCron)+" ")]}}])}),a("el-table-column",{attrs:{align:"center",prop:"created_at",label:"创建时间",width:"200"},scopedSlots:t._u([{key:"default",fn:function(e){var n=e.row;return[a("span",[t._v(t._s(n.created_at))])]}}])}),a("el-table-column",{attrs:{align:"center",label:"操作"},scopedSlots:t._u([{key:"default",fn:function(e){var n=e.row;return[16&n.optionbutton?a("el-button",{attrs:{type:"info",size:"small"},on:{click:function(e){return t.showDetail(n)}}},[t._v("详情")]):t._e(),a("router-link",{attrs:{to:{name:"editTask",params:{id:n.id}},tag:"span"}},[8&n.optionbutton?a("el-button",{attrs:{type:"warning",size:"small",icon:"edit"}},[t._v("修改")]):t._e()],1),4&n.optionbutton?a("el-button",{attrs:{type:"danger",size:"small"},on:{click:function(e){return t.stop(n.id)}}},[t._v("停止")]):t._e(),2&n.optionbutton?a("el-button",{attrs:{type:"success",size:"small"},on:{click:function(e){return t.start(n.id)}}},[t._v("启动")]):t._e(),1&n.optionbutton?a("el-button",{attrs:{type:"danger",size:"small"},on:{click:function(e){return t.del(n.id)}}},[t._v("删除")]):t._e()]}}])})],1),a("div",{staticClass:"pagination-container fr"},[a("el-pagination",{attrs:{background:"","current-page":t.currentPage,"page-sizes":[10,20,30,50],"page-size":t.size,layout:"total, sizes, prev, pager, next, jumper",total:t.total},on:{"size-change":t.handleSizeChange,"current-change":t.handleCurrentChange}})],1),a("el-dialog",{attrs:{title:"任务详情",visible:t.showDetails,width:"60%",center:""},on:{"update:visible":function(e){t.showDetails=e}}},[a("el-row",{attrs:{border:"",fit:"","highlight-current-row":""}},[a("el-form",{staticClass:"task-desc-form",attrs:{"label-position":"right","label-width":"150px","label-suffix":":"}},[a("el-col",{attrs:{span:24}},[a("el-form-item",{attrs:{label:"编号"}},[a("span",[t._v(t._s(t.details.id))])]),a("el-form-item",{attrs:{label:"名称"}},[a("span",[t._v(t._s(t.details.task_name))])]),a("el-form-item",{attrs:{label:"任务规则名"}},[a("span",[t._v(t._s(t.details.task_rule_name))])]),a("el-form-item",{attrs:{label:"任务描述"}},[a("span",[t._v(t._s(t.details.task_desc))])]),a("el-form-item",{attrs:{label:"定时任务"}},[a("span",[t._v(t._s(t.details.isCron))])]),a("el-form-item",{attrs:{label:"定时执行"}},[a("span",[t._v(t._s(t.details.cron_spec))])]),a("el-form-item",{attrs:{label:"代理列表"}},[a("span",[t._v(t._s(t.details.proxy_urls))])]),a("el-form-item",{attrs:{label:"用户代理"}},[a("span",[t._v(t._s(t.details.opt_user_agent))])]),a("el-form-item",{attrs:{label:"爬虫最大深度"}},[a("span",[t._v(t._s(t.details.opt_max_depth))])]),a("el-form-item",{attrs:{label:"允许访问的域名"}},[a("span",[t._v(t._s(t.details.opt_allowed_domains))])]),a("el-form-item",{attrs:{label:"URL过滤"}},[a("span",[t._v(t._s(t.details.opt_url_filters))])]),a("el-form-item",{attrs:{label:"最大body值"}},[a("span",[t._v(t._s(t.details.opt_max_body_size))])]),a("el-form-item",{attrs:{label:"请求超时时间"}},[a("span",[t._v(t._s(t.details.opt_request_timeout))])]),a("el-form-item",{attrs:{label:"频率限制"}},[a("span",[t._v(t._s(t.details.limit_enable))])]),a("el-form-item",{attrs:{label:"域名glob匹配"}},[a("span",[t._v(t._s(t.details.limit_domain_glob))])]),a("el-form-item",{attrs:{label:"延迟"}},[a("span",[t._v(t._s(t.details.limit_delay))])]),a("el-form-item",{attrs:{label:"随机延迟"}},[a("span",[t._v(t._s(t.details.limit_random_delay))])]),a("el-form-item",{attrs:{label:"请求并发度"}},[a("span",[t._v(t._s(t.details.limit_parallelism))])]),a("el-form-item",{attrs:{label:"状态"}},[a("span",[t._v(t._s(t._f("statusFilter")(t.details.status)))])]),a("el-form-item",{attrs:{label:"运行次数"}},[a("span",[t._v(t._s(t.details.counts))])]),a("el-form-item",{attrs:{label:"创建时间"}},[a("span",[t._v(t._s(t.details.created_at))])])],1)],1)],1),a("span",{staticClass:"dialog-footer",attrs:{slot:"footer"},slot:"footer"},[a("el-button",{on:{click:function(e){t.showDetails=!1}}},[t._v("Close")])],1)],1)],1)},s=[];a("a4d3"),a("e01a"),a("d3b7"),a("d28b"),a("3ca3"),a("ddb0"),a("fb6a"),a("b0c0"),a("a630"),a("ac1f"),a("00b4");function r(t,e){(null==e||e>t.length)&&(e=t.length);for(var a=0,n=new Array(e);a<e;a++)n[a]=t[a];return n}function i(t,e){if(t){if("string"===typeof t)return r(t,e);var a=Object.prototype.toString.call(t).slice(8,-1);return"Object"===a&&t.constructor&&(a=t.constructor.name),"Map"===a||"Set"===a?Array.from(t):"Arguments"===a||/^(?:Ui|I)nt(?:8|16|32)(?:Clamped)?Array$/.test(a)?r(t,e):void 0}}function o(t,e){var a="undefined"!==typeof Symbol&&t[Symbol.iterator]||t["@@iterator"];if(!a){if(Array.isArray(t)||(a=i(t))||e&&t&&"number"===typeof t.length){a&&(t=a);var n=0,s=function(){};return{s:s,n:function(){return n>=t.length?{done:!0}:{done:!1,value:t[n++]}},e:function(t){throw t},f:s}}throw new TypeError("Invalid attempt to iterate non-iterable instance.\nIn order to be iterable, non-array objects must have a [Symbol.iterator]() method.")}var r,o=!0,l=!1;return{s:function(){a=a.call(t)},n:function(){var t=a.next();return o=t.done,t},e:function(t){l=!0,r=t},f:function(){try{o||null==a["return"]||a["return"]()}finally{if(l)throw r}}}}var l=a("b199"),c={filters:{statusFilter:function(t){var e={0:"初始状态",1:"运行中",2:"停止",3:"完成",4:"异常终止"};return e[t]}},data:function(){return{list:null,listLoading:!0,currentPage:1,total:0,size:10,details:{},showDetails:!1}},created:function(){this.fetchData()},methods:{onRefresh:function(){this.fetchData()},fetchData:function(){var t=this;this.listLoading=!0,Object(l["b"])({offset:(this.currentPage-1)*this.size,size:this.size}).then((function(e){t.list=e.data.list;var a,n=o(t.list);try{for(n.s();!(a=n.n()).done;){var s=a.value;switch(s.isCron="No",s.cron_spec&&(s.isCron="Yes"),s.status){case 0:s.optionbutton=27;break;case 1:s.optionbutton=20;break;case 2:s.optionbutton=25,s.cron_spec&&(s.optionbutton=27);break;case 3:s.optionbutton=20;break;case 4:s.optionbutton=20;break;default:s.optionbutton=27;break}}}catch(r){n.e(r)}finally{n.f()}t.total=e.data.total,t.listLoading=!1})).catch((function(){t.listLoading=!1}))},showDetail:function(t){this.showDetails=!0,this.details=t},stop:function(t){var e=this;this.$confirm("该操作将会停止该任务，是否继续操作？","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((function(){e.listLoading=!0,Object(l["g"])(t).then((function(){e.fetchData(),e.$message.success("操作成功")})).catch((function(){e.$message.error("操作失败")}))})).catch((function(t){console.log(t)}))},start:function(t){var e=this;this.listLoading=!0,Object(l["f"])(t).then((function(){e.fetchData(),e.$message.success("操作成功")})).catch((function(){e.$message.error("操作失败")})).catch((function(t){console.log(t)}))},del:function(t){var e=this;this.listLoading=!0,Object(l["a"])(t).then((function(){e.fetchData(),e.$message.success("操作成功")})).catch((function(){e.$message.error("操作失败")})).catch((function(t){console.log(t)}))},handleSizeChange:function(t){this.size=t,this.fetchData()},handleCurrentChange:function(t){this.currentPage=t,this.fetchData()}}},u=c,f=(a("1fe5"),a("2877")),d=Object(f["a"])(u,n,s,!1,null,"357a60ac",null);e["default"]=d.exports},b199:function(t,e,a){"use strict";a.d(e,"b",(function(){return s})),a.d(e,"d",(function(){return r})),a.d(e,"c",(function(){return i})),a.d(e,"g",(function(){return o})),a.d(e,"f",(function(){return l})),a.d(e,"e",(function(){return c})),a.d(e,"a",(function(){return u}));var n=a("b775");function s(t){return Object(n["a"])({url:"/webapi/tasks",method:"get",params:t})}function r(t){return Object(n["a"])({url:"/webapi/tasks/"+t,method:"get"})}function i(){return Object(n["a"])({url:"/webapi/tasks/rules",method:"get"})}function o(t){return Object(n["a"])({url:"/webapi/tasks/"+t+"/stop",method:"get"})}function l(t){return Object(n["a"])({url:"/webapi/tasks/"+t+"/start",method:"get"})}function c(t){return Object(n["a"])({url:"/webapi/tasks",method:"post",data:t})}function u(t){return Object(n["a"])({url:"/webapi/tasks/"+t+"/delete",method:"get"})}}}]);