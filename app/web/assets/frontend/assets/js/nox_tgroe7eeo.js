import{r as p,j as t,f as y,v as c,F as k,B as r}from"./nox_ophygXhXg.js";import{a as i}from"./nox_Xrtokpkyt.js";import{T as d}from"./nox_yppAogorA.js";import{S}from"./nox_77oQgerAX.js";import{m}from"./nox_kkeQohppr.js";import"./nox_tre7yXyrQ.js";import"./nox_pk7goeQyA.js";import"./nox_kgkgAgQr7.js";import"./nox_krekpQe7r.js";import"./nox_7AoggoAXX.js";import"./nox_eA7p7A7ey.js";function P(e){const[s,n]=p.useState(!1),{sender:g}=i(),{sender:u}=i(),{sender:l,pending:h}=i(),f=()=>{l({method:"put",path:"tokopedia/mapper/map",payload:e.list.map(a=>({shopee_id:a.shopeeCategoryId,tokopedia_id:a.tokopediaCategoryIds.map(Number)[a.tokopediaCategoryIds.length-1]}))})},x=()=>{l({method:"put",path:"tokopedia/mapper/map",payload:e.list.map(a=>({shopee_id:a.shopeeCategoryId,tokopedia_id:0}))},{onSuccess(){m.success("Update data success!")},onError(){m.error("Update data error!")}}),e.namespace&&e.initEffect(e.namespace,e.listCategoryTokopedia)},C=()=>{!s&&e.namespace&&(n(!0),g({method:"put",path:"tokopedia/mapper/autosuggest",params:{collection:e.namespace}}))};return p.useEffect(()=>{const a=setInterval(()=>{if(!s)return e.namespace&&!s&&e.initEffect(e.namespace,e.listCategoryTokopedia),clearInterval(a);u({method:"get",path:"tokopedia/mapper/autosuggest"},{onSuccess:o=>{(o==null?void 0:o.status)==="STOPPED"&&n(!1)},onError:o=>{n(!1),console.log(o)}})},1e3);return()=>clearInterval(a)},[s]),t.jsx(y,{size:"small",title:t.jsx(d.Text,{children:"Map Category From Shopee to Tokopedia"}),children:t.jsxs(c,{style:{justifyContent:"space-between",alignItems:"end"},children:[t.jsxs(k,{style:{rowGap:"5px"},children:[t.jsx(d.Text,{children:"Collections :"}),t.jsx(S,{style:{width:"300px"},placeholder:"Choose Collection",value:e.namespace,onChange:e.onChangeNamespace,options:e.namespaces})]}),t.jsxs(c,{style:{rowGap:"5px",justifyContent:"flex-end"},children:[t.jsx(r,{loading:s,disabled:s,onClick:C,children:"Use Suggest"}),t.jsx(r,{style:{backgroundColor:"#005246"},type:"primary",onClick:x,children:"Reset All"}),t.jsx(r,{loading:h,onClick:f,type:"primary",children:"Save Mapping"})]})]})})}export{P as default};
