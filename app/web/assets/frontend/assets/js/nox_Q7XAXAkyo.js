import{r as i,b as e,C as s,j as l,J as n,F as m}from"./nox_Ao7tXQoge.js";import{I as c}from"./nox_ygho7kAgX.js";import{T as d}from"./nox_kXtpeggoX.js";import{D as p}from"./nox_7o7kpkooQ.js";import{M as f}from"./nox_p7og7hphA.js";import"./nox_prAehkXhg.js";import"./nox_QXg7pXkkp.js";import"./nox_hggQyoy7A.js";import"./nox_oXreoyXoy.js";import"./nox_yAoXQgok7.js";import"./nox_AQgQXghgt.js";function T(t){const[r,o]=i.useState("");return e(f,{width:390,footer:!1,closable:!1,centered:!0,...t,children:e(s,{title:"File Target Name",size:"small",type:"inner",children:l(m,{style:{rowGap:5},children:[e(c,{value:r,onChange:a=>o(a.target.value),addonAfter:".csv",placeholder:"Boleh dikosongi"}),e(d.Text,{children:"Default filename: cekbot.csv"}),e(p,{dashed:!0,style:{marginBlock:4}}),e(n,{onClick:()=>{r?t.onFinish(r):t.onFinish("cekbot.csv")},type:"primary",children:"Run"})]})})})}export{T as default};