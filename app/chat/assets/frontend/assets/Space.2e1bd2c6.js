import{ap as B,d as C,J as G,K as S,L as R,c as j,M as E,by as L,aW as x,Q as v}from"./index.db4b08e0.js";import{f as M}from"./index.eb8c481c.js";import{g as I}from"./get-slot.65c4337d.js";const P={gapSmall:"4px 8px",gapMedium:"8px 12px",gapLarge:"12px 16px"},T=()=>P,W={name:"Space",self:T},A=W;let h;const N=()=>{if(!B)return!0;if(h===void 0){const e=document.createElement("div");e.style.display="flex",e.style.flexDirection="column",e.style.rowGap="1px",e.appendChild(document.createElement("div")),e.appendChild(document.createElement("div")),document.body.appendChild(e);const n=e.scrollHeight===1;return document.body.removeChild(e),h=n}return h},O=Object.assign(Object.assign({},S.props),{align:String,justify:{type:String,default:"start"},inline:Boolean,vertical:Boolean,size:{type:[String,Number,Array],default:"medium"},wrapItem:{type:Boolean,default:!0},itemStyle:[String,Object],wrap:{type:Boolean,default:!0},internalUseGap:{type:Boolean,default:void 0}}),J=C({name:"Space",props:O,setup(e){const{mergedClsPrefixRef:n,mergedRtlRef:u}=G(e),a=S("Space","-space",void 0,A,e,n),g=R("Space",u,n);return{useGap:N(),rtlEnabled:g,mergedClsPrefix:n,margin:j(()=>{const{size:t}=e;if(Array.isArray(t))return{horizontal:t[0],vertical:t[1]};if(typeof t=="number")return{horizontal:t,vertical:t};const{self:{[E("gap",t)]:f}}=a.value,{row:l,col:p}=L(f);return{horizontal:x(p),vertical:x(l)}})}},render(){const{vertical:e,align:n,inline:u,justify:a,itemStyle:g,margin:t,wrap:f,mergedClsPrefix:l,rtlEnabled:p,useGap:o,wrapItem:b,internalUseGap:w}=this,c=M(I(this));if(!c.length)return null;const y=`${t.horizontal}px`,m=`${t.horizontal/2}px`,$=`${t.vertical}px`,i=`${t.vertical/2}px`,s=c.length-1,d=a.startsWith("space-");return v("div",{role:"none",class:[`${l}-space`,p&&`${l}-space--rtl`],style:{display:u?"inline-flex":"flex",flexDirection:e?"column":"row",justifyContent:["start","end"].includes(a)?"flex-"+a:a,flexWrap:!f||e?"nowrap":"wrap",marginTop:o||e?"":`-${i}`,marginBottom:o||e?"":`-${i}`,alignItems:n,gap:o?`${t.vertical}px ${t.horizontal}px`:""}},!b&&(o||w)?c:c.map((z,r)=>v("div",{role:"none",style:[g,{maxWidth:"100%"},o?"":e?{marginBottom:r!==s?$:""}:p?{marginLeft:d?a==="space-between"&&r===s?"":m:r!==s?y:"",marginRight:d?a==="space-between"&&r===0?"":m:"",paddingTop:i,paddingBottom:i}:{marginRight:d?a==="space-between"&&r===s?"":m:r!==s?y:"",marginLeft:d?a==="space-between"&&r===0?"":m:"",paddingTop:i,paddingBottom:i}]},z)))}});export{J as N};
