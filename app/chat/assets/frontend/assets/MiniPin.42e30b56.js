import{d,o as i,g as l,h as n,b as o,j as p,a as s,T as c,u as _,k as f,p as u,m as r,q as y}from"./index.db4b08e0.js";const m=e=>(u("data-v-e17678f5"),e=e(),r(),e),b={style:{overflow:"hidden"},class:"is-relative"},h=m(()=>o("i",{style:{visibility:"hidden"},class:"bi fs-11 bi-pin-fill"},null,-1)),v={key:0,style:{position:"absolute",top:"0",left:"0",color:"#555"},class:"bi fs-11 bi-pin-fill"},k={key:1,style:{position:"absolute",top:"0",left:"0"},class:"fs-11 bi bi-pin"},B=d({__name:"MiniPin",props:{pinned:{type:Boolean}},emits:["update:pinned"],setup(e,{emit:a}){return(x,t)=>(i(),l(_(f),{onClick:t[0]||(t[0]=I=>a("update:pinned",!e.pinned)),size:"tiny",type:"default",secondary:""},{default:n(()=>[o("i",b,[h,p(c,{name:e.pinned?"unpinned":"pinned"},{default:n(()=>[e.pinned?(i(),s("i",v)):(i(),s("i",k))]),_:1},8,["name"])])]),_:1}))}});const w=y(B,[["__scopeId","data-v-e17678f5"]]);export{w as M};