import{r as u,b as a,j as s,E as p,k as C,J as g,F as _,C as k}from"./nox_Ao7tXQoge.js";import{u as S}from"./nox_ekpgkQhpQ.js";import{B as b}from"./nox_ktAgoQpA7.js";import{S as h}from"./nox_7teyQ7Akr.js";import{D as w}from"./nox_7o7kpkooQ.js";import{I as x}from"./nox_ygho7kAgX.js";import{m as l}from"./nox_yeeXokpeX.js";import"./nox_oXreoyXoy.js";import"./nox_prAehkXhg.js";import"./nox_A7tek7epA.js";import"./nox_yAoXQgok7.js";import"./nox_QXg7pXkkp.js";import"./nox_khpkttAoh.js";import"./nox_hggQyoy7A.js";import"./nox_AQgQXghgt.js";function q(t){var c;const[o,n]=u.useState(),[f,r]=u.useState(t.item.EtalaseName||void 0),{sender:d}=S();function y(e){o||d({method:"put",path:"tokopedia/etalase_map/update",payload:[{category_id:t.item.tokopedia_id,etalase_name:e,ID:0}]},{onSuccess(){r(e),l.success(`Updated as ${e}`)}})}function m(){o&&t.item.tokopedia_id&&d({method:"put",path:"tokopedia/etalase_map/update",payload:[{category_id:t.item.tokopedia_id,etalase_name:o,ID:0}]},{onError(e){l.error(JSON.stringify(e))},onSuccess(){var e;l.success(`Updated as ${o}`),(e=t.refetchFn)==null||e.call(t,{method:"get",path:"tokopedia/etalase_map/list_etalase"},{onSuccess(){r(o)}})}})}return a(k,{size:"small",type:"inner",style:t.style,children:s(_,{style:{justifyContent:"space-between",height:"100%"},children:[s(p,{style:{justifyContent:"space-between",flex:1},children:[a(b,{items:t.item.ShopeeCategoryName.map(e=>({title:e})),separator:">"}),a("div",{style:{width:"20px"}}),s("span",{style:{flexShrink:0},children:[a("strong",{children:t.item.product_count||0})," Product"]})]}),a(h,{placeholder:t.item.tokopedia_id?"Choose etalase":"Check category mapping",disabled:!t.item.tokopedia_id,dropdownRender:e=>s(C,{children:[e,a(w,{style:{marginBlock:4}}),s(p,{style:{columnGap:5},children:[a(x,{value:o,onKeyDown:i=>{i.code=="Enter"&&m()},onChange:i=>n(i.target.value),placeholder:"Create new etalase..."}),a(g,{disabled:!o,onClick:m,type:"primary",children:"Add"})]})]}),value:f,onChange:e=>y(e),onClick:()=>n(void 0),children:(c=t.etalases)==null?void 0:c.sort((e,i)=>e.label.toLowerCase()<i.label.toLowerCase()?-1:0).map(e=>a(h.Option,{children:e.label},e.label))})]})})}export{q as default};
