import{r as i,A as se,_ as ae,aA as rt,K as ee,x as it,aP as lt,g as st,f as te,ab as at,C as we,aC as Te,a9 as $e,aI as Z,ad as le,Y as Re,au as ct,aJ as dt}from"./nox_rghtApo7t.js";import{u as ut,C as pt}from"./nox_ykrXhoohh.js";import{T as ft}from"./nox_heQgkAkrk.js";import{i as gt}from"./nox_yyXeoApor.js";var mt={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M832 64H296c-4.4 0-8 3.6-8 8v56c0 4.4 3.6 8 8 8h496v688c0 4.4 3.6 8 8 8h56c4.4 0 8-3.6 8-8V96c0-17.7-14.3-32-32-32zM704 192H192c-17.7 0-32 14.3-32 32v530.7c0 8.5 3.4 16.6 9.4 22.6l173.3 173.3c2.2 2.2 4.7 4 7.4 5.5v1.9h4.2c3.5 1.3 7.2 2 11 2H704c17.7 0 32-14.3 32-32V224c0-17.7-14.3-32-32-32zM350 856.2L263.9 770H350v86.2zM664 888H414V746c0-22.1-17.9-40-40-40H232V264h432v624z"}}]},name:"copy",theme:"outlined"};const yt=mt;var bt=function(t,o){return i.createElement(se,ae({},t,{ref:o,icon:yt}))};const vt=i.forwardRef(bt);var ht={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M257.7 752c2 0 4-.2 6-.5L431.9 722c2-.4 3.9-1.3 5.3-2.8l423.9-423.9a9.96 9.96 0 000-14.1L694.9 114.9c-1.9-1.9-4.4-2.9-7.1-2.9s-5.2 1-7.1 2.9L256.8 538.8c-1.5 1.5-2.4 3.3-2.8 5.3l-29.5 168.2a33.5 33.5 0 009.4 29.8c6.6 6.4 14.9 9.9 23.8 9.9zm67.4-174.4L687.8 215l73.3 73.3-362.7 362.6-88.9 15.7 15.6-89zM880 836H144c-17.7 0-32 14.3-32 32v36c0 4.4 3.6 8 8 8h784c4.4 0 8-3.6 8-8v-36c0-17.7-14.3-32-32-32z"}}]},name:"edit",theme:"outlined"};const Et=ht;var St=function(t,o){return i.createElement(se,ae({},t,{ref:o,icon:Et}))};const xt=i.forwardRef(St);var Ot={icon:{tag:"svg",attrs:{viewBox:"64 64 896 896",focusable:"false"},children:[{tag:"path",attrs:{d:"M864 170h-60c-4.4 0-8 3.6-8 8v518H310v-73c0-6.7-7.8-10.5-13-6.3l-141.9 112a8 8 0 000 12.6l141.9 112c5.3 4.2 13 .4 13-6.3v-75h498c35.3 0 64-28.7 64-64V178c0-4.4-3.6-8-8-8z"}}]},name:"enter",theme:"outlined"};const Ct=Ot;var wt=function(t,o){return i.createElement(se,ae({},t,{ref:o,icon:Ct}))};const Tt=i.forwardRef(wt),je=e=>({color:e.colorLink,textDecoration:"none",outline:"none",cursor:"pointer",transition:`color ${e.motionDurationSlow}`,"&:focus, &:hover":{color:e.colorLinkHover},"&:active":{color:e.colorLinkActive}});var Ie=function(t){if(rt()&&window.document.documentElement){var o=Array.isArray(t)?t:[t],n=window.document.documentElement;return o.some(function(r){return r in n.style})}return!1},$t=function(t,o){if(!Ie(t))return!1;var n=document.createElement("div"),r=n.style[t];return n.style[t]=o,n.style[t]!==r};function ve(e,t){return!Array.isArray(e)&&t!==void 0?$t(e,t):Ie(e)}var Rt=globalThis&&globalThis.__rest||function(e,t){var o={};for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&t.indexOf(n)<0&&(o[n]=e[n]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var r=0,n=Object.getOwnPropertySymbols(e);r<n.length;r++)t.indexOf(n[r])<0&&Object.prototype.propertyIsEnumerable.call(e,n[r])&&(o[n[r]]=e[n[r]]);return o};const jt={border:0,background:"transparent",padding:0,lineHeight:"inherit",display:"inline-block"},It=i.forwardRef((e,t)=>{const o=g=>{const{keyCode:s}=g;s===ee.ENTER&&g.preventDefault()},n=g=>{const{keyCode:s}=g,{onClick:y}=e;s===ee.ENTER&&y&&y()},{style:r,noStyle:p,disabled:m}=e,c=Rt(e,["style","noStyle","disabled"]);let d={};return p||(d=Object.assign({},jt)),m&&(d.pointerEvents="none"),d=Object.assign(Object.assign({},d),r),i.createElement("div",Object.assign({role:"button",tabIndex:0,ref:t},c,{onKeyDown:o,onKeyUp:n,style:d}))}),he=It;var Lt=function(){var e=document.getSelection();if(!e.rangeCount)return function(){};for(var t=document.activeElement,o=[],n=0;n<e.rangeCount;n++)o.push(e.getRangeAt(n));switch(t.tagName.toUpperCase()){case"INPUT":case"TEXTAREA":t.blur();break;default:t=null;break}return e.removeAllRanges(),function(){e.type==="Caret"&&e.removeAllRanges(),e.rangeCount||o.forEach(function(r){e.addRange(r)}),t&&t.focus()}},Pt=Lt,Ee={"text/plain":"Text","text/html":"Url",default:"Text"},Dt="Copy to clipboard: #{key}, Enter";function _t(e){var t=(/mac os x/i.test(navigator.userAgent)?"⌘":"Ctrl")+"+C";return e.replace(/#{\s*key\s*}/g,t)}function kt(e,t){var o,n,r,p,m,c,d=!1;t||(t={}),o=t.debug||!1;try{r=Pt(),p=document.createRange(),m=document.getSelection(),c=document.createElement("span"),c.textContent=e,c.ariaHidden="true",c.style.all="unset",c.style.position="fixed",c.style.top=0,c.style.clip="rect(0, 0, 0, 0)",c.style.whiteSpace="pre",c.style.webkitUserSelect="text",c.style.MozUserSelect="text",c.style.msUserSelect="text",c.style.userSelect="text",c.addEventListener("copy",function(s){if(s.stopPropagation(),t.format)if(s.preventDefault(),typeof s.clipboardData>"u"){o&&console.warn("unable to use e.clipboardData"),o&&console.warn("trying IE specific stuff"),window.clipboardData.clearData();var y=Ee[t.format]||Ee.default;window.clipboardData.setData(y,e)}else s.clipboardData.clearData(),s.clipboardData.setData(t.format,e);t.onCopy&&(s.preventDefault(),t.onCopy(s.clipboardData))}),document.body.appendChild(c),p.selectNodeContents(c),m.addRange(p);var g=document.execCommand("copy");if(!g)throw new Error("copy command was unsuccessful");d=!0}catch(s){o&&console.error("unable to copy using execCommand: ",s),o&&console.warn("trying IE specific stuff");try{window.clipboardData.setData(t.format||"text",e),t.onCopy&&t.onCopy(window.clipboardData),d=!0}catch(y){o&&console.error("unable to copy using clipboardData: ",y),o&&console.error("falling back to prompt"),n=_t("message"in t?t.message:Dt),window.prompt(n,e)}}finally{m&&(typeof m.removeRange=="function"?m.removeRange(p):m.removeAllRanges()),c&&document.body.removeChild(c),r()}return d}var zt=kt;const Ht=it(zt),Nt=(e,t,o,n)=>{const{sizeMarginHeadingVerticalEnd:r,fontWeightStrong:p}=n;return{marginBottom:r,color:o,fontWeight:p,fontSize:e,lineHeight:t}},Mt=e=>{const t=[1,2,3,4,5],o={};return t.forEach(n=>{o[`
      h${n}&,
      div&-h${n},
      div&-h${n} > textarea,
      h${n}
    `]=Nt(e[`fontSizeHeading${n}`],e[`lineHeightHeading${n}`],e.colorTextHeading,e)}),o},At=e=>{const{componentCls:t}=e;return{"a&, a":Object.assign(Object.assign({},je(e)),{textDecoration:e.linkDecoration,"&:active, &:hover":{textDecoration:e.linkHoverDecoration},[`&[disabled], &${t}-disabled`]:{color:e.colorTextDisabled,cursor:"not-allowed","&:active, &:hover":{color:e.colorTextDisabled},"&:active":{pointerEvents:"none"}}})}},Bt=e=>({code:{margin:"0 0.2em",paddingInline:"0.4em",paddingBlock:"0.2em 0.1em",fontSize:"85%",fontFamily:e.fontFamilyCode,background:"rgba(150, 150, 150, 0.1)",border:"1px solid rgba(100, 100, 100, 0.2)",borderRadius:3},kbd:{margin:"0 0.2em",paddingInline:"0.4em",paddingBlock:"0.15em 0.1em",fontSize:"90%",fontFamily:e.fontFamilyCode,background:"rgba(150, 150, 150, 0.06)",border:"1px solid rgba(100, 100, 100, 0.2)",borderBottomWidth:2,borderRadius:3},mark:{padding:0,backgroundColor:lt[2]},"u, ins":{textDecoration:"underline",textDecorationSkipInk:"auto"},"s, del":{textDecoration:"line-through"},strong:{fontWeight:600},"ul, ol":{marginInline:0,marginBlock:"0 1em",padding:0,li:{marginInline:"20px 0",marginBlock:0,paddingInline:"4px 0",paddingBlock:0}},ul:{listStyleType:"circle",ul:{listStyleType:"disc"}},ol:{listStyleType:"decimal"},"pre, blockquote":{margin:"1em 0"},pre:{padding:"0.4em 0.6em",whiteSpace:"pre-wrap",wordWrap:"break-word",background:"rgba(150, 150, 150, 0.1)",border:"1px solid rgba(100, 100, 100, 0.2)",borderRadius:3,fontFamily:e.fontFamilyCode,code:{display:"inline",margin:0,padding:0,fontSize:"inherit",fontFamily:"inherit",background:"transparent",border:0}},blockquote:{paddingInline:"0.6em 0",paddingBlock:0,borderInlineStart:"4px solid rgba(100, 100, 100, 0.2)",opacity:.85}}),Wt=e=>{const{componentCls:t}=e,n=gt(e).inputPaddingVertical+1;return{"&-edit-content":{position:"relative","div&":{insetInlineStart:-e.paddingSM,marginTop:-n,marginBottom:`calc(1em - ${n}px)`},[`${t}-edit-content-confirm`]:{position:"absolute",insetInlineEnd:e.marginXS+2,insetBlockEnd:e.marginXS,color:e.colorTextDescription,fontWeight:"normal",fontSize:e.fontSize,fontStyle:"normal",pointerEvents:"none"},textarea:{margin:"0!important",MozTransition:"none",height:"1em"}}}},Kt=e=>({"&-copy-success":{[`
    &,
    &:hover,
    &:focus`]:{color:e.colorSuccess}}}),Ut=()=>({[`
  a&-ellipsis,
  span&-ellipsis
  `]:{display:"inline-block",maxWidth:"100%"},"&-single-line":{whiteSpace:"nowrap"},"&-ellipsis-single-line":{overflow:"hidden",textOverflow:"ellipsis","a&, span&":{verticalAlign:"bottom"}},"&-ellipsis-multiple-line":{display:"-webkit-box",overflow:"hidden",WebkitLineClamp:3,WebkitBoxOrient:"vertical"}}),Vt=e=>{const{componentCls:t,sizeMarginHeadingVerticalStart:o}=e;return{[t]:Object.assign(Object.assign(Object.assign(Object.assign(Object.assign(Object.assign(Object.assign(Object.assign(Object.assign({color:e.colorText,wordBreak:"break-word",lineHeight:e.lineHeight,[`&${t}-secondary`]:{color:e.colorTextDescription},[`&${t}-success`]:{color:e.colorSuccess},[`&${t}-warning`]:{color:e.colorWarning},[`&${t}-danger`]:{color:e.colorError,"a&:active, a&:focus":{color:e.colorErrorActive},"a&:hover":{color:e.colorErrorHover}},[`&${t}-disabled`]:{color:e.colorTextDisabled,cursor:"not-allowed",userSelect:"none"},[`
        div&,
        p
      `]:{marginBottom:"1em"}},Mt(e)),{[`
      & + h1${t},
      & + h2${t},
      & + h3${t},
      & + h4${t},
      & + h5${t}
      `]:{marginTop:o},[`
      div,
      ul,
      li,
      p,
      h1,
      h2,
      h3,
      h4,
      h5`]:{[`
        + h1,
        + h2,
        + h3,
        + h4,
        + h5
        `]:{marginTop:o}}}),Bt(e)),At(e)),{[`
        ${t}-expand,
        ${t}-edit,
        ${t}-copy
      `]:Object.assign(Object.assign({},je(e)),{marginInlineStart:e.marginXXS})}),Wt(e)),Kt(e)),Ut()),{"&-rtl":{direction:"rtl"}})}},Le=st("Typography",e=>[Vt(e)],{sizeMarginHeadingVerticalStart:"1.2em",sizeMarginHeadingVerticalEnd:"0.5em"}),Ft=e=>{let{prefixCls:t,"aria-label":o,className:n,style:r,direction:p,maxLength:m,autoSize:c=!0,value:d,onSave:g,onCancel:s,onEnd:y,component:b,enterIcon:$=i.createElement(Tt,null)}=e;const S=i.useRef(null),C=i.useRef(!1),_=i.useRef(),[I,w]=i.useState(d);i.useEffect(()=>{w(d)},[d]),i.useEffect(()=>{if(S.current&&S.current.resizableTextArea){const{textArea:O}=S.current.resizableTextArea;O.focus();const{length:R}=O.value;O.setSelectionRange(R,R)}},[]);const u=O=>{let{target:R}=O;w(R.value.replace(/[\n\r]/g,""))},M=()=>{C.current=!0},A=()=>{C.current=!1},v=O=>{let{keyCode:R}=O;C.current||(_.current=R)},W=()=>{g(I.trim())},h=O=>{let{keyCode:R,ctrlKey:oe,altKey:V,metaKey:z,shiftKey:K}=O;_.current===R&&!C.current&&!oe&&!V&&!z&&!K&&(R===ee.ENTER?(W(),y==null||y()):R===ee.ESC&&s())},f=()=>{W()},T=b?`${t}-${b}`:"",[B,k]=Le(t),D=te(t,`${t}-edit-content`,{[`${t}-rtl`]:p==="rtl"},n,T,k);return B(i.createElement("div",{className:D,style:r},i.createElement(ft,{ref:S,maxLength:m,value:I,onChange:u,onKeyDown:v,onKeyUp:h,onCompositionStart:M,onCompositionEnd:A,onBlur:f,"aria-label":o,rows:1,autoSize:c}),$!==null?at($,{className:`${t}-edit-content-confirm`}):null))},Xt=Ft;function re(e,t){return i.useMemo(()=>{const o=!!e;return[o,Object.assign(Object.assign({},t),o&&typeof e=="object"?e:null)]},[e])}const Jt=(e,t)=>{const o=i.useRef(!1);i.useEffect(()=>{o.current?e():o.current=!0},t)},qt=Jt;var Gt=globalThis&&globalThis.__rest||function(e,t){var o={};for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&t.indexOf(n)<0&&(o[n]=e[n]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var r=0,n=Object.getOwnPropertySymbols(e);r<n.length;r++)t.indexOf(n[r])<0&&Object.prototype.propertyIsEnumerable.call(e,n[r])&&(o[n[r]]=e[n[r]]);return o};const Yt=i.forwardRef((e,t)=>{var{prefixCls:o,component:n="article",className:r,rootClassName:p,setContentRef:m,children:c,direction:d}=e,g=Gt(e,["prefixCls","component","className","rootClassName","setContentRef","children","direction"]);const{getPrefixCls:s,direction:y}=i.useContext(we),b=d??y;let $=t;m&&($=Te(t,m));const S=s("typography",o),[C,_]=Le(S),I=te(S,{[`${S}-rtl`]:b==="rtl"},r,p,_);return C(i.createElement(n,Object.assign({className:I,ref:$},g),c))}),Pe=Yt;function De(e){const t=typeof e;return t==="string"||t==="number"}function Qt(e){let t=0;return e.forEach(o=>{De(o)?t+=String(o).length:t+=1}),t}function Se(e,t){let o=0;const n=[];for(let r=0;r<e.length;r+=1){if(o===t)return n;const p=e[r],c=De(p)?String(p).length:1,d=o+c;if(d>t){const g=t-o;return n.push(String(p).slice(0,g)),n}n.push(p),o=d}return e}const Zt=0,Y=1,xe=2,ie=3,Oe=4,en=e=>{let{enabledMeasure:t,children:o,text:n,width:r,fontSize:p,rows:m,onEllipsis:c}=e;const[[d,g,s],y]=i.useState([0,0,0]),[b,$]=i.useState(Zt),[S,C]=i.useState(0),_=i.useRef(null),I=i.useRef(null),w=i.useMemo(()=>$e(n),[n]),u=i.useMemo(()=>Qt(w),[w]),M=i.useMemo(()=>!t||b!==ie?o(w,!1):o(Se(w,g),g<u),[t,b,o,w,g,u]);Z(()=>{t&&r&&p&&u&&($(Y),y([0,Math.ceil(u/2),u]))},[t,r,p,n,u,m]),Z(()=>{var h;b===Y&&C(((h=_.current)===null||h===void 0?void 0:h.offsetHeight)||0)},[b]),Z(()=>{var h,f;if(S){if(b===Y){const T=((h=I.current)===null||h===void 0?void 0:h.offsetHeight)||0,B=m*S;T<=B?($(Oe),c(!1)):$(xe)}else if(b===xe)if(d!==s){const T=((f=I.current)===null||f===void 0?void 0:f.offsetHeight)||0,B=m*S;let k=d,D=s;d===s-1?D=d:T<=B?k=g:D=g;const O=Math.ceil((k+D)/2);y([k,O,D])}else $(ie),c(!0)}},[b,d,s,m,S]);const A={width:r,whiteSpace:"normal",margin:0,padding:0},v=(h,f,T)=>i.createElement("span",{"aria-hidden":!0,ref:f,style:Object.assign({position:"fixed",display:"block",left:0,top:0,zIndex:-9999,visibility:"hidden",pointerEvents:"none",fontSize:Math.floor(p/2)*2},T)},h),W=(h,f)=>{const T=Se(w,h);return v(o(T,!0),f,A)};return i.createElement(i.Fragment,null,M,t&&b!==ie&&b!==Oe&&i.createElement(i.Fragment,null,v("lg",_,{wordBreak:"keep-all",whiteSpace:"nowrap"}),b===Y?v(o(w,!1),I,A):W(g,I)))},tn=en,nn=e=>{let{enabledEllipsis:t,isEllipsis:o,children:n,tooltipProps:r}=e;return!(r!=null&&r.title)||!t?n:i.createElement(le,Object.assign({open:o?void 0:!1},r),n)},on=nn;var rn=globalThis&&globalThis.__rest||function(e,t){var o={};for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&t.indexOf(n)<0&&(o[n]=e[n]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var r=0,n=Object.getOwnPropertySymbols(e);r<n.length;r++)t.indexOf(n[r])<0&&Object.prototype.propertyIsEnumerable.call(e,n[r])&&(o[n[r]]=e[n[r]]);return o};function ln(e,t){let{mark:o,code:n,underline:r,delete:p,strong:m,keyboard:c,italic:d}=e,g=t;function s(y,b){b&&(g=i.createElement(y,{},g))}return s("strong",m),s("u",r),s("del",p),s("code",n),s("mark",o),s("kbd",c),s("i",d),g}function Q(e,t,o){return e===!0||e===void 0?t:e||o&&t}function Ce(e){return e===!1?[!1,!1]:Array.isArray(e)?e:[e]}const sn="...",an=i.forwardRef((e,t)=>{var o,n,r;const{prefixCls:p,className:m,style:c,type:d,disabled:g,children:s,ellipsis:y,editable:b,copyable:$,component:S,title:C}=e,_=rn(e,["prefixCls","className","style","type","disabled","children","ellipsis","editable","copyable","component","title"]),{getPrefixCls:I,direction:w}=i.useContext(we),[u]=ut("Text"),M=i.useRef(null),A=i.useRef(null),v=I("typography",p),W=Re(_,["mark","code","delete","underline","strong","keyboard","italic"]),[h,f]=re(b),[T,B]=ct(!1,{value:f.editing}),{triggerType:k=["icon"]}=f,D=l=>{var a;l&&((a=f.onStart)===null||a===void 0||a.call(f)),B(l)};qt(()=>{var l;T||(l=A.current)===null||l===void 0||l.focus()},[T]);const O=l=>{l==null||l.preventDefault(),D(!0)},R=l=>{var a;(a=f.onChange)===null||a===void 0||a.call(f,l),D(!1)},oe=()=>{var l;(l=f.onCancel)===null||l===void 0||l.call(f),D(!1)},[V,z]=re($),[K,ce]=i.useState(!1),de=i.useRef(),ue={};z.format&&(ue.format=z.format);const pe=()=>{window.clearTimeout(de.current)},_e=l=>{var a;l==null||l.preventDefault(),l==null||l.stopPropagation(),Ht(z.text||String(s)||"",ue),ce(!0),pe(),de.current=window.setTimeout(()=>{ce(!1)},3e3),(a=z.onCopy)===null||a===void 0||a.call(z,l)};i.useEffect(()=>pe,[]);const[fe,ke]=i.useState(!1),[ge,ze]=i.useState(!1),[He,Ne]=i.useState(!1),[me,Me]=i.useState(!1),[ye,Ae]=i.useState(!1),[Be,We]=i.useState(!0),[H,E]=re(y,{expandable:!1}),L=H&&!He,{rows:U=1}=E,X=i.useMemo(()=>!L||E.suffix!==void 0||E.onEllipsis||E.expandable||h||V,[L,E,h,V]);Z(()=>{H&&!X&&(ke(ve("webkitLineClamp")),ze(ve("textOverflow")))},[X,H]);const P=i.useMemo(()=>X?!1:U===1?ge:fe,[X,ge,fe]),be=L&&(P?ye:me),Ke=L&&U===1&&P,J=L&&U>1&&P,Ue=l=>{var a;Ne(!0),(a=E.onExpand)===null||a===void 0||a.call(E,l)},[Ve,Fe]=i.useState(0),[Xe,Je]=i.useState(0),qe=(l,a)=>{let{offsetWidth:x}=l;var j;Fe(x),Je(parseInt((j=window.getComputedStyle)===null||j===void 0?void 0:j.call(window,a).fontSize,10)||0)},Ge=l=>{var a;Me(l),me!==l&&((a=E.onEllipsis)===null||a===void 0||a.call(E,l))};i.useEffect(()=>{const l=M.current;if(H&&P&&l){const a=J?l.offsetHeight<l.scrollHeight:l.offsetWidth<l.scrollWidth;ye!==a&&Ae(a)}},[H,P,s,J,Be]),i.useEffect(()=>{const l=M.current;if(typeof IntersectionObserver>"u"||!l||!P||!L)return;const a=new IntersectionObserver(()=>{We(!!l.offsetParent)});return a.observe(l),()=>{a.disconnect()}},[P,L]);let N={};E.tooltip===!0?N={title:(o=f.text)!==null&&o!==void 0?o:s}:i.isValidElement(E.tooltip)?N={title:E.tooltip}:typeof E.tooltip=="object"?N=Object.assign({title:(n=f.text)!==null&&n!==void 0?n:s},E.tooltip):N={title:E.tooltip};const q=i.useMemo(()=>{const l=a=>["string","number"].includes(typeof a);if(!(!H||P)){if(l(f.text))return f.text;if(l(s))return s;if(l(C))return C;if(l(N.title))return N.title}},[H,P,C,N.title,be]);if(T)return i.createElement(Xt,{value:(r=f.text)!==null&&r!==void 0?r:typeof s=="string"?s:"",onSave:R,onCancel:oe,onEnd:f.onEnd,prefixCls:v,className:m,style:c,direction:w,component:S,maxLength:f.maxLength,autoSize:f.autoSize,enterIcon:f.enterIcon});const Ye=()=>{const{expandable:l,symbol:a}=E;if(!l)return null;let x;return a?x=a:x=u==null?void 0:u.expand,i.createElement("a",{key:"expand",className:`${v}-expand`,onClick:Ue,"aria-label":u==null?void 0:u.expand},x)},Qe=()=>{if(!h)return;const{icon:l,tooltip:a}=f,x=$e(a)[0]||(u==null?void 0:u.edit),j=typeof x=="string"?x:"";return k.includes("icon")?i.createElement(le,{key:"edit",title:a===!1?"":x},i.createElement(he,{ref:A,className:`${v}-edit`,onClick:O,"aria-label":j},l||i.createElement(xt,{role:"button"}))):null},Ze=()=>{if(!V)return;const{tooltips:l,icon:a}=z,x=Ce(l),j=Ce(a),G=K?Q(x[1],u==null?void 0:u.copied):Q(x[0],u==null?void 0:u.copy),nt=K?u==null?void 0:u.copied:u==null?void 0:u.copy,ot=typeof G=="string"?G:nt;return i.createElement(le,{key:"copy",title:G},i.createElement(he,{className:te(`${v}-copy`,K&&`${v}-copy-success`),onClick:_e,"aria-label":ot},K?Q(j[1],i.createElement(pt,null),!0):Q(j[0],i.createElement(vt,null),!0)))},et=l=>[l&&Ye(),Qe(),Ze()],tt=l=>[l&&i.createElement("span",{"aria-hidden":!0,key:"ellipsis"},sn),E.suffix,et(l)];return i.createElement(dt,{onResize:qe,disabled:!L||P},l=>i.createElement(on,{tooltipProps:N,enabledEllipsis:L,isEllipsis:be},i.createElement(Pe,Object.assign({className:te({[`${v}-${d}`]:d,[`${v}-disabled`]:g,[`${v}-ellipsis`]:H,[`${v}-single-line`]:L&&U===1,[`${v}-ellipsis-single-line`]:Ke,[`${v}-ellipsis-multiple-line`]:J},m),prefixCls:p,style:Object.assign(Object.assign({},c),{WebkitLineClamp:J?U:void 0}),component:S,ref:Te(l,M,t),direction:w,onClick:k.includes("text")?O:void 0,"aria-label":q==null?void 0:q.toString(),title:C},W),i.createElement(tn,{enabledMeasure:L&&!P,text:s,rows:U,width:Ve,fontSize:Xe,onEllipsis:Ge},(a,x)=>{let j=a;return a.length&&x&&q&&(j=i.createElement("span",{key:"show-content","aria-hidden":!0},j)),ln(e,i.createElement(i.Fragment,null,j,tt(x)))}))))}),ne=an;var cn=globalThis&&globalThis.__rest||function(e,t){var o={};for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&t.indexOf(n)<0&&(o[n]=e[n]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var r=0,n=Object.getOwnPropertySymbols(e);r<n.length;r++)t.indexOf(n[r])<0&&Object.prototype.propertyIsEnumerable.call(e,n[r])&&(o[n[r]]=e[n[r]]);return o};const dn=i.forwardRef((e,t)=>{var{ellipsis:o,rel:n}=e,r=cn(e,["ellipsis","rel"]);const p=Object.assign(Object.assign({},r),{rel:n===void 0&&r.target==="_blank"?"noopener noreferrer":n});return delete p.navigate,i.createElement(ne,Object.assign({},p,{ref:t,ellipsis:!!o,component:"a"}))}),un=dn,pn=i.forwardRef((e,t)=>i.createElement(ne,Object.assign({ref:t},e,{component:"div"}))),fn=pn;var gn=globalThis&&globalThis.__rest||function(e,t){var o={};for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&t.indexOf(n)<0&&(o[n]=e[n]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var r=0,n=Object.getOwnPropertySymbols(e);r<n.length;r++)t.indexOf(n[r])<0&&Object.prototype.propertyIsEnumerable.call(e,n[r])&&(o[n[r]]=e[n[r]]);return o};const mn=(e,t)=>{var{ellipsis:o}=e,n=gn(e,["ellipsis"]);const r=i.useMemo(()=>o&&typeof o=="object"?Re(o,["expandable","rows"]):o,[o]);return i.createElement(ne,Object.assign({ref:t},n,{ellipsis:r,component:"span"}))},yn=i.forwardRef(mn);var bn=globalThis&&globalThis.__rest||function(e,t){var o={};for(var n in e)Object.prototype.hasOwnProperty.call(e,n)&&t.indexOf(n)<0&&(o[n]=e[n]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var r=0,n=Object.getOwnPropertySymbols(e);r<n.length;r++)t.indexOf(n[r])<0&&Object.prototype.propertyIsEnumerable.call(e,n[r])&&(o[n[r]]=e[n[r]]);return o};const vn=[1,2,3,4,5],hn=i.forwardRef((e,t)=>{const{level:o=1}=e,n=bn(e,["level"]);let r;return vn.includes(o)?r=`h${o}`:r="h1",i.createElement(ne,Object.assign({ref:t},n,{component:r}))}),En=hn,F=Pe;F.Text=yn;F.Link=un;F.Title=En;F.Paragraph=fn;const wn=F;export{vt as C,wn as T};
