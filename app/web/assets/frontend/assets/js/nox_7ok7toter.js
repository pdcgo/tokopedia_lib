import{r as i,g as K,m as Q,c as T,d as F,e as M,f as U,D as Y,h as j,W as Z,T as ee,u as oe,i as te,j as L,b as E,C as W,F as re,R as H,_ as V}from"./nox_yApeAXQky.js";import{F as ne,p as ie}from"./nox_yeQo7eteh.js";import{C as ae}from"./nox_y7rgtoAyh.js";import{D as le}from"./nox_7r7rghpoe.js";const X=i.createContext(null),de=X.Provider,q=i.createContext(null),se=q.Provider,ce=o=>{const{componentCls:r,antCls:n}=o,t=`${r}-group`;return{[t]:Object.assign(Object.assign({},T(o)),{display:"inline-block",fontSize:0,[`&${t}-rtl`]:{direction:"rtl"},[`${n}-badge ${n}-badge-count`]:{zIndex:1},[`> ${n}-badge:not(:first-child) > ${n}-button-wrapper`]:{borderInlineStart:"none"}})}},ue=o=>{const{componentCls:r,wrapperMarginInlineEnd:n,colorPrimary:t,radioSize:e,motionDurationSlow:c,motionDurationMid:d,motionEaseInOutCirc:u,colorBgContainer:l,colorBorder:S,lineWidth:p,dotSize:k,colorBgContainerDisabled:h,colorTextDisabled:v,paddingXS:R,dotColorDisabled:O,lineType:C,radioDotDisabledSize:x,wireframe:f,colorWhite:w}=o,b=`${r}-inner`;return{[`${r}-wrapper`]:Object.assign(Object.assign({},T(o)),{display:"inline-flex",alignItems:"baseline",marginInlineStart:0,marginInlineEnd:n,cursor:"pointer",[`&${r}-wrapper-rtl`]:{direction:"rtl"},"&-disabled":{cursor:"not-allowed",color:o.colorTextDisabled},"&::after":{display:"inline-block",width:0,overflow:"hidden",content:'"\\a0"'},[`${r}-checked::after`]:{position:"absolute",insetBlockStart:0,insetInlineStart:0,width:"100%",height:"100%",border:`${p}px ${C} ${t}`,borderRadius:"50%",visibility:"hidden",content:'""'},[r]:Object.assign(Object.assign({},T(o)),{position:"relative",display:"inline-block",outline:"none",cursor:"pointer",alignSelf:"center",borderRadius:"50%"}),[`${r}-wrapper:hover &,
        &:hover ${b}`]:{borderColor:t},[`${r}-input:focus-visible + ${b}`]:Object.assign({},F(o)),[`${r}:hover::after, ${r}-wrapper:hover &::after`]:{visibility:"visible"},[`${r}-inner`]:{"&::after":{boxSizing:"border-box",position:"absolute",insetBlockStart:"50%",insetInlineStart:"50%",display:"block",width:e,height:e,marginBlockStart:e/-2,marginInlineStart:e/-2,backgroundColor:f?t:w,borderBlockStart:0,borderInlineStart:0,borderRadius:e,transform:"scale(0)",opacity:0,transition:`all ${c} ${u}`,content:'""'},boxSizing:"border-box",position:"relative",insetBlockStart:0,insetInlineStart:0,display:"block",width:e,height:e,backgroundColor:l,borderColor:S,borderStyle:"solid",borderWidth:p,borderRadius:"50%",transition:`all ${d}`},[`${r}-input`]:{position:"absolute",insetBlockStart:0,insetInlineEnd:0,insetBlockEnd:0,insetInlineStart:0,width:0,height:0,padding:0,margin:0,zIndex:1,cursor:"pointer",opacity:0},[`${r}-checked`]:{[b]:{borderColor:t,backgroundColor:f?l:t,"&::after":{transform:`scale(${k/e})`,opacity:1,transition:`all ${c} ${u}`}}},[`${r}-disabled`]:{cursor:"not-allowed",[b]:{backgroundColor:h,borderColor:S,cursor:"not-allowed","&::after":{backgroundColor:O}},[`${r}-input`]:{cursor:"not-allowed"},[`${r}-disabled + span`]:{color:v,cursor:"not-allowed"},[`&${r}-checked`]:{[b]:{"&::after":{transform:`scale(${x/e})`}}}},[`span${r} + *`]:{paddingInlineStart:R,paddingInlineEnd:R}})}},be=o=>{const{buttonColor:r,controlHeight:n,componentCls:t,lineWidth:e,lineType:c,colorBorder:d,motionDurationSlow:u,motionDurationMid:l,buttonPaddingInline:S,fontSize:p,buttonBg:k,fontSizeLG:h,controlHeightLG:v,controlHeightSM:R,paddingXS:O,borderRadius:C,borderRadiusSM:x,borderRadiusLG:f,buttonCheckedBg:w,buttonSolidCheckedColor:b,colorTextDisabled:g,colorBgContainerDisabled:I,buttonCheckedBgDisabled:B,buttonCheckedColorDisabled:s,colorPrimary:m,colorPrimaryHover:$,colorPrimaryActive:y}=o;return{[`${t}-button-wrapper`]:{position:"relative",display:"inline-block",height:n,margin:0,paddingInline:S,paddingBlock:0,color:r,fontSize:p,lineHeight:`${n-e*2}px`,background:k,border:`${e}px ${c} ${d}`,borderBlockStartWidth:e+.02,borderInlineStartWidth:0,borderInlineEndWidth:e,cursor:"pointer",transition:[`color ${l}`,`background ${l}`,`box-shadow ${l}`].join(","),a:{color:r},[`> ${t}-button`]:{position:"absolute",insetBlockStart:0,insetInlineStart:0,zIndex:-1,width:"100%",height:"100%"},"&:not(:first-child)":{"&::before":{position:"absolute",insetBlockStart:-e,insetInlineStart:-e,display:"block",boxSizing:"content-box",width:1,height:"100%",paddingBlock:e,paddingInline:0,backgroundColor:d,transition:`background-color ${u}`,content:'""'}},"&:first-child":{borderInlineStart:`${e}px ${c} ${d}`,borderStartStartRadius:C,borderEndStartRadius:C},"&:last-child":{borderStartEndRadius:C,borderEndEndRadius:C},"&:first-child:last-child":{borderRadius:C},[`${t}-group-large &`]:{height:v,fontSize:h,lineHeight:`${v-e*2}px`,"&:first-child":{borderStartStartRadius:f,borderEndStartRadius:f},"&:last-child":{borderStartEndRadius:f,borderEndEndRadius:f}},[`${t}-group-small &`]:{height:R,paddingInline:O-e,paddingBlock:0,lineHeight:`${R-e*2}px`,"&:first-child":{borderStartStartRadius:x,borderEndStartRadius:x},"&:last-child":{borderStartEndRadius:x,borderEndEndRadius:x}},"&:hover":{position:"relative",color:m},"&:has(:focus-visible)":Object.assign({},F(o)),[`${t}-inner, input[type='checkbox'], input[type='radio']`]:{width:0,height:0,opacity:0,pointerEvents:"none"},[`&-checked:not(${t}-button-wrapper-disabled)`]:{zIndex:1,color:m,background:w,borderColor:m,"&::before":{backgroundColor:m},"&:first-child":{borderColor:m},"&:hover":{color:$,borderColor:$,"&::before":{backgroundColor:$}},"&:active":{color:y,borderColor:y,"&::before":{backgroundColor:y}}},[`${t}-group-solid &-checked:not(${t}-button-wrapper-disabled)`]:{color:b,background:m,borderColor:m,"&:hover":{color:b,background:$,borderColor:$},"&:active":{color:b,background:y,borderColor:y}},"&-disabled":{color:g,backgroundColor:I,borderColor:d,cursor:"not-allowed","&:first-child, &:hover":{color:g,backgroundColor:I,borderColor:d}},[`&-disabled${t}-button-wrapper-checked`]:{color:s,backgroundColor:B,borderColor:d,boxShadow:"none"}}}},A=o=>o-4*2,J=K("Radio",o=>{const{controlOutline:r,controlOutlineWidth:n,radioSize:t}=o,e=`0 0 0 ${n}px ${r}`,c=e,d=A(t),u=Q(o,{radioDotDisabledSize:d,radioFocusShadow:e,radioButtonFocusShadow:c});return[ce(u),ue(u),be(u)]},o=>{const{wireframe:r,padding:n,marginXS:t,lineWidth:e,fontSizeLG:c,colorText:d,colorBgContainer:u,colorTextDisabled:l,controlItemBgActiveDisabled:S,colorTextLightSolid:p}=o,k=4,h=c,v=r?A(h):h-(k+e)*2;return{radioSize:h,dotSize:v,dotColorDisabled:l,buttonSolidCheckedColor:p,buttonBg:u,buttonCheckedBg:u,buttonColor:d,buttonCheckedBgDisabled:S,buttonCheckedColorDisabled:l,buttonPaddingInline:n-e,wrapperMarginInlineEnd:t}});var ge=globalThis&&globalThis.__rest||function(o,r){var n={};for(var t in o)Object.prototype.hasOwnProperty.call(o,t)&&r.indexOf(t)<0&&(n[t]=o[t]);if(o!=null&&typeof Object.getOwnPropertySymbols=="function")for(var e=0,t=Object.getOwnPropertySymbols(o);e<t.length;e++)r.indexOf(t[e])<0&&Object.prototype.propertyIsEnumerable.call(o,t[e])&&(n[t[e]]=o[t[e]]);return n};const pe=(o,r)=>{var n,t;const e=i.useContext(X),c=i.useContext(q),{getPrefixCls:d,direction:u,radio:l}=i.useContext(M),S=i.useRef(null),p=U(r,S),{isFormItemInput:k}=i.useContext(ne),h=y=>{var a,P;(a=o.onChange)===null||a===void 0||a.call(o,y),(P=e==null?void 0:e.onChange)===null||P===void 0||P.call(e,y)},{prefixCls:v,className:R,rootClassName:O,children:C,style:x}=o,f=ge(o,["prefixCls","className","rootClassName","children","style"]),w=d("radio",v),b=((e==null?void 0:e.optionType)||c)==="button",g=b?`${w}-button`:w,[I,B]=J(w),s=Object.assign({},f),m=i.useContext(Y);e&&(s.name=e.name,s.onChange=h,s.checked=o.value===e.value,s.disabled=(n=s.disabled)!==null&&n!==void 0?n:e.disabled),s.disabled=(t=s.disabled)!==null&&t!==void 0?t:m;const $=j(`${g}-wrapper`,{[`${g}-wrapper-checked`]:s.checked,[`${g}-wrapper-disabled`]:s.disabled,[`${g}-wrapper-rtl`]:u==="rtl",[`${g}-wrapper-in-form-item`]:k},l==null?void 0:l.className,R,O,B);return I(i.createElement(Z,{component:"Radio",disabled:s.disabled},i.createElement("label",{className:$,style:Object.assign(Object.assign({},l==null?void 0:l.style),x),onMouseEnter:o.onMouseEnter,onMouseLeave:o.onMouseLeave},i.createElement(ae,Object.assign({},s,{className:j(s.className,!b&&ee),type:"radio",prefixCls:g,ref:p})),C!==void 0?i.createElement("span",null,C):null)))},he=i.forwardRef(pe),D=he,Ce=i.forwardRef((o,r)=>{const{getPrefixCls:n,direction:t}=i.useContext(M),[e,c]=oe(o.defaultValue,{value:o.value}),d=a=>{const P=e,N=a.target.value;"value"in o||c(N);const{onChange:G}=o;G&&N!==P&&G(a)},{prefixCls:u,className:l,rootClassName:S,options:p,buttonStyle:k="outline",disabled:h,children:v,size:R,style:O,id:C,onMouseEnter:x,onMouseLeave:f,onFocus:w,onBlur:b}=o,g=n("radio",u),I=`${g}-group`,[B,s]=J(g);let m=v;p&&p.length>0&&(m=p.map(a=>typeof a=="string"||typeof a=="number"?i.createElement(D,{key:a.toString(),prefixCls:g,disabled:h,value:a,checked:e===a},a):i.createElement(D,{key:`radio-group-value-options-${a.value}`,prefixCls:g,disabled:a.disabled||h,value:a.value,checked:e===a.value,title:a.title,style:a.style},a.label)));const $=te(R),y=j(I,`${I}-${k}`,{[`${I}-${$}`]:$,[`${I}-rtl`]:t==="rtl"},l,S,s);return B(i.createElement("div",Object.assign({},ie(o,{aria:!0,data:!0}),{className:y,style:O,onMouseEnter:x,onMouseLeave:f,onFocus:w,onBlur:b,id:C,ref:r}),i.createElement(de,{value:{onChange:d,value:e,disabled:o.disabled,name:o.name,optionType:o.optionType}},m)))}),me=i.memo(Ce);var fe=globalThis&&globalThis.__rest||function(o,r){var n={};for(var t in o)Object.prototype.hasOwnProperty.call(o,t)&&r.indexOf(t)<0&&(n[t]=o[t]);if(o!=null&&typeof Object.getOwnPropertySymbols=="function")for(var e=0,t=Object.getOwnPropertySymbols(o);e<t.length;e++)r.indexOf(t[e])<0&&Object.prototype.propertyIsEnumerable.call(o,t[e])&&(n[t[e]]=o[t[e]]);return n};const Se=(o,r)=>{const{getPrefixCls:n}=i.useContext(M),{prefixCls:t}=o,e=fe(o,["prefixCls"]),c=n("radio",t);return i.createElement(se,{value:"button"},i.createElement(D,Object.assign({prefixCls:c},e,{type:"radio",ref:r})))},ve=i.forwardRef(Se),_=D;_.Button=ve;_.Group=me;_.__ANT_RADIO=!0;const z=_,xe=H.lazy(()=>V(()=>import("./nox_ryet7rp7k.js"),["./nox_ryet7rp7k.js","./nox_yApeAXQky.js","..\\css\\nox_rtXQpkheA.css","./nox_eeptQph7h.js","./nox_Ahhk7ghgk.js","./nox_eXtp7kApQ.js","./nox_yeQo7eteh.js","./nox_ghhrtoQko.js","./nox_hhh7hektt.js","./nox_Xyo77tytQ.js","./nox_ryegreegg.js","./nox_Aeprprtpe.js"],import.meta.url)),$e=H.lazy(()=>V(()=>import("./nox_QtkeAQX7h.js"),["./nox_QtkeAQX7h.js","./nox_yApeAXQky.js","..\\css\\nox_rtXQpkheA.css","./nox_eeptQph7h.js","./nox_hAArkAorp.js","./nox_krkhkpoAp.js","./nox_ppyrX7keA.js","./nox_Xyo77tytQ.js","./nox_hhh7hektt.js","./nox_yeQo7eteh.js","./nox_Ahhk7ghgk.js","./nox_eXtp7kApQ.js","./nox_ghhrtoQko.js","./nox_Xgg7yetQQ.js","./nox_etekoAAtA.js","./nox_geyAQpp7p.js","./nox_Aeprprtpe.js","./nox_ryegreegg.js"],import.meta.url));function Ie(){const[o,r]=i.useState("common_menu");return L(re,{children:[L(z.Group,{value:o,onChange:n=>r(n.target.value),children:[E(z.Button,{value:"common_menu",children:"Common Menu"}),E(z.Button,{value:"deleter",children:"Tokopedia Product Deleter"})]}),E(le,{dashed:!0,style:{marginBlock:2}}),o=="common_menu"&&E(i.Suspense,{fallback:E(W,{loading:!0}),children:E(xe,{})}),o=="deleter"&&E(i.Suspense,{fallback:E(W,{loading:!0}),children:E($e,{})})]})}export{Ie as default};