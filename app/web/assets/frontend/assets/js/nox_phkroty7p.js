import{b as i,j as o,aa as m,E as h,F as t,C as g}from"./nox_Ao7tXQoge.js";import{S as n}from"./nox_7teyQ7Akr.js";import{C as u}from"./nox_gyryphkr7.js";import{C as b,T as a}from"./nox_kXtpeggoX.js";import{m as f}from"./nox_yeeXokpeX.js";import{I as p}from"./nox_ygho7kAgX.js";import{I as x}from"./nox_krre7oyg7.js";import{F,D as w}from"./nox_r7oote7kX.js";import"./nox_A7tek7epA.js";import"./nox_prAehkXhg.js";import"./nox_yAoXQgok7.js";import"./nox_oXreoyXoy.js";import"./nox_QXg7pXkkp.js";import"./nox_ekpgkQhpQ.js";import"./nox_krerotQoe.js";import"./nox_AQgQXghgt.js";import"./nox_tQkkgoAAo.js";import"./nox_ygt7XrtAe.js";import"./nox_khpkttAoh.js";import"./nox_ep7prgrXp.js";import"./nox_XrXorhXtg.js";import"./nox_7o7kpkooQ.js";import"./nox_hggQyoy7A.js";const S=e=>{const{mode:r,collections:d,manualCollections:l,value:y,...C}=e,s=r==="tokopedia_manual"?l:d;console.log(r);const c=s.find(v=>v.label===y);return i(n,{value:c==null?void 0:c.value,options:[{value:"",label:"Choose Collection",disabled:!0},...s],placeholder:"Choose Collection Data",...C})},k=S;function V(e){var r,d;return i(g,{title:o(u,{checked:e.profile.isChecked,onChange:l=>{e.updateSingleProfileFn(e.profile.id,{isChecked:l.target.checked})},style:{userSelect:"none"},children:[e.number+". ",e.profile.id]}),hoverable:!0,size:"small",type:"inner",actions:[i(m,{title:"Copy",placement:"bottom",showArrow:!1,children:i(b,{style:{color:"#FFA559"},rev:"copy",onClick:()=>{f.success(`Copied profile: ${e.profile.id}`),e.copyProfileFn(e.profile)}},"copy")}),i(m,{title:"Paste",placement:"bottom",showArrow:!1,children:i(F,{style:{color:"#FFA559"},rev:"paste",onClick:()=>{e.clipboard&&(f.info(`Paste from: ${e.clipboard.id}`),e.updateSingleProfileFn(e.profile.id,{limitUpload:e.clipboard.limitUpload,markupName:e.clipboard.markupName,spinName:e.clipboard.spinName,colName:e.clipboard.colName}))}},"paste")}),i(m,{title:"Remove",placement:"bottom",showArrow:!1,children:i(w,{style:{color:"#FFA559"},onClick:()=>e.deleter({method:"post",path:"tokopedia/akun/delete",payload:{usernames:[e.profile.id]}}),rev:"delete"},"delete")})],children:o(t,{children:[o(h,{style:{width:"100%"},children:[o(t,{style:{flex:1},children:[o(t,{style:{rowGap:"5px"},children:[i(a.Text,{type:"secondary",children:"Username :"}),i(p,{value:e.profile.emailOrUsername,onChange:l=>e.updateSingleProfileFn(e.profile.id,{emailOrUsername:l.target.value}),placeholder:"username"})]}),o(t,{style:{rowGap:"5px"},children:[i(a.Text,{type:"secondary",children:"Password :"}),i(p.Password,{value:e.profile.password,onChange:l=>e.updateSingleProfileFn(e.profile.id,{password:l.target.value}),placeholder:"⁎⁎⁎⁎⁎⁎⁎⁎"})]}),o(t,{style:{rowGap:"5px"},children:[i(a.Text,{type:"secondary",children:"Upload Limit :"}),i(x,{value:e.profile.limitUpload,onChange:l=>e.updateSingleProfileFn(e.profile.id,{limitUpload:l||0}),placeholder:"1000",style:{width:"100%"}})]}),i("div",{})]}),o(t,{style:{flex:1},children:[o(t,{style:{rowGap:"5px"},children:[i(a.Text,{type:"secondary",children:"Markup :"}),o(n,{value:e.profile.markupName,onChange:l=>e.updateSingleProfileFn(e.profile.id,{markupName:l}),placeholder:"Choose Markup Data",children:[i(n.Option,{disabled:!0,value:"",children:"Markup Select"}),(r=e.markups)==null?void 0:r.map(l=>i(n.Option,{value:l.value,children:l.label},l.value))]})]}),o(t,{style:{rowGap:"5px"},children:[i(a.Text,{type:"secondary",children:"Spin :"}),o(n,{value:e.profile.spinName,onChange:l=>e.updateSingleProfileFn(e.profile.id,{spinName:l}),placeholder:"Choose Spin Data",children:[i(n.Option,{disabled:!0,value:"",children:"Spin Select"}),(d=e.spins)==null?void 0:d.map(l=>i(n.Option,{value:l.value,children:l.label},l.value))]})]}),o(t,{style:{rowGap:"5px"},children:[i(a.Text,{type:"secondary",children:"Collection :"}),i(k,{value:e.profile.colName,mode:e.mode,collections:e.collections,manualCollections:e.manualCollections,onChange:l=>e.updateSingleProfileFn(e.profile.id,{colName:l})})]})]})]}),o(h,{style:{justifyContent:"space-between",width:"100%"},children:[i(u,{checked:e.profile.isActive,onChange:l=>e.updateSingleProfileFn(e.profile.id,{isActive:l.target.checked}),style:{userSelect:"none"},children:"Active"}),o(a.Text,{children:["Product Uploaded Count:"," ",e.profile.productCount||0]})]})]})})}export{V as default};