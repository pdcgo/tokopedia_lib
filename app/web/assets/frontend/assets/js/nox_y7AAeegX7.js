import{r as c,A as o,_ as d,k as t,l as r,n as i,B as n}from"./nox_rghtApo7t.js";import{C as s,F as h,D as u}from"./nox_QhhoAotXp.js";import{I as v}from"./nox_yQgAApgoh.js";import{C as m}from"./nox_ykrXhoohh.js";import"./nox_yyXeoApor.js";import"./nox_pAAteykoX.js";import"./nox_heQgkAkrk.js";import"./nox_Q7hhk7k7e.js";var x={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M893.3 293.3L730.7 130.7c-7.5-7.5-16.7-13-26.7-16V112H144c-17.7 0-32 14.3-32 32v736c0 17.7 14.3 32 32 32h736c17.7 0 32-14.3 32-32V338.5c0-17-6.7-33.2-18.7-45.2zM384 184h256v104H384V184zm456 656H184V184h136v136c0 17.7 14.3 32 32 32h320c17.7 0 32-14.3 32-32V205.8l136 136V840zM512 442c-79.5 0-144 64.5-144 144s64.5 144 144 144 144-64.5 144-144-64.5-144-144-144zm0 224c-44.2 0-80-35.8-80-80s35.8-80 80-80 80 35.8 80 80-35.8 80-80 80z"}}]},name:"save",theme:"outlined"};const f=x;var C=function(l,a){return c.createElement(o,d({},l,{ref:a,icon:f}))};const g=c.forwardRef(C);var S={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M400 317.7h73.9V656c0 4.4 3.6 8 8 8h60c4.4 0 8-3.6 8-8V317.7H624c6.7 0 10.4-7.7 6.3-12.9L518.3 163a8 8 0 00-12.6 0l-112 141.7c-4.1 5.3-.4 13 6.3 13zM878 626h-60c-4.4 0-8 3.6-8 8v154H214V634c0-4.4-3.6-8-8-8h-60c-4.4 0-8 3.6-8 8v198c0 17.7 14.3 32 32 32h684c17.7 0 32-14.3 32-32V634c0-4.4-3.6-8-8-8z"}}]},name:"upload",theme:"outlined"};const j=S;var k=function(l,a){return c.createElement(o,d({},l,{ref:a,icon:j}))};const O=c.forwardRef(k);function M(e){return t.jsx(r,{size:"small",title:"Setting Tokopedia Upload",children:t.jsxs(i,{style:{justifyContent:"space-between",alignItems:"center"},children:[t.jsx(s,{checked:e.checkedAll,onChange:l=>{var a;(a=e.onChangeCheckedAll)==null||a.call(e,l.target.checked)},children:"Select All"}),t.jsxs(i,{style:{flex:1},children:[t.jsx(v,{allowClear:!0,placeholder:"Search Profile...",style:{flex:1},value:e.nameQuery,onChange:l=>{var a;return(a=e.onChangeNameQuery)==null?void 0:a.call(e,l.target.value)}}),t.jsx(n,{onClick:e.onClickPasteAll,icon:t.jsx(h,{rev:"paste"}),disabled:e.disablePasteAll,children:"Paste All"}),t.jsx(n,{onClick:e.onClickSetActive,icon:t.jsx(m,{rev:"active"}),children:"Active All"}),t.jsx(n,{disabled:!0,icon:t.jsx(u,{rev:"remove"}),children:"Remove"}),t.jsx(n,{type:"primary",icon:t.jsx(g,{rev:"save"}),onClick:e.onClickSave,loading:e.loadingSave,style:{backgroundColor:"#C2418D",boxShadow:"none",color:"#fff"},children:"Save"}),t.jsx(n,{type:"primary",icon:t.jsx(O,{rev:"upload"}),style:{boxShadow:"none"},onClick:e.onClickStartUpload,loading:e.loadingStartUpload,children:"Start Upload"})]})]})})}export{M as default};
