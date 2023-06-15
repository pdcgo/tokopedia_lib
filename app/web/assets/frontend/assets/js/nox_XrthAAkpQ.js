import{r as a,E as _,_ as K,f as j,p as oe,K as se,y as M,q as A,a0 as ae,a1 as re,a2 as ie,g as ce,m as le,I as z,e as ue,C as Q,a3 as fe,j as de,U as me,S as ge,V as W}from"./nox_rghtApo7t.js";import{C as ve,a as pe,E as ye}from"./nox_Q7hhk7k7e.js";import{I as Ce}from"./nox_Qeyoptr7e.js";var xe=a.forwardRef(function(e,n){var s=e.prefixCls,t=e.style,o=e.className,i=e.duration,f=i===void 0?4.5:i,x=e.eventKey,l=e.content,u=e.closable,g=e.closeIcon,m=g===void 0?"x":g,d=e.props,p=e.onClick,C=e.onNoticeClose,b=e.times,N=a.useState(!1),$=_(N,2),h=$[0],S=$[1],O=function(){C(x)},E=function(y){(y.key==="Enter"||y.code==="Enter"||y.keyCode===se.ENTER)&&O()};a.useEffect(function(){if(!h&&f>0){var r=setTimeout(function(){O()},f*1e3);return function(){clearTimeout(r)}}},[f,h,b]);var c="".concat(s,"-notice");return a.createElement("div",K({},d,{ref:n,className:j(c,o,oe({},"".concat(c,"-closable"),u)),style:t,onMouseEnter:function(){S(!0)},onMouseLeave:function(){S(!1)},onClick:p}),a.createElement("div",{className:"".concat(c,"-content")},l),u&&a.createElement("a",{tabIndex:0,className:"".concat(c,"-close"),onKeyDown:E,onClick:function(y){y.preventDefault(),y.stopPropagation(),O()}},m))});const X=xe;var be=a.forwardRef(function(e,n){var s=e.prefixCls,t=s===void 0?"rc-notification":s,o=e.container,i=e.motion,f=e.maxCount,x=e.className,l=e.style,u=e.onAllRemoved,g=a.useState([]),m=_(g,2),d=m[0],p=m[1],C=function(r){var y,v=d.find(function(P){return P.key===r});v==null||(y=v.onClose)===null||y===void 0||y.call(v),p(function(P){return P.filter(function(w){return w.key!==r})})};a.useImperativeHandle(n,function(){return{open:function(r){p(function(y){var v=M(y),P=v.findIndex(function(H){return H.key===r.key}),w=A({},r);if(P>=0){var R;w.times=(((R=y[P])===null||R===void 0?void 0:R.times)||0)+1,v[P]=w}else w.times=0,v.push(w);return f>0&&v.length>f&&(v=v.slice(-f)),v})},close:function(r){C(r)},destroy:function(){p([])}}});var b=a.useState({}),N=_(b,2),$=N[0],h=N[1];a.useEffect(function(){var c={};d.forEach(function(r){var y=r.placement,v=y===void 0?"topRight":y;v&&(c[v]=c[v]||[],c[v].push(r))}),Object.keys($).forEach(function(r){c[r]=c[r]||[]}),h(c)},[d]);var S=function(r){h(function(y){var v=A({},y),P=v[r]||[];return P.length||delete v[r],v})},O=a.useRef(!1);if(a.useEffect(function(){Object.keys($).length>0?O.current=!0:O.current&&(u==null||u(),O.current=!1)},[$]),!o)return null;var E=Object.keys($);return ae.createPortal(a.createElement(a.Fragment,null,E.map(function(c){var r=$[c],y=r.map(function(P){return{config:P,key:P.key}}),v=typeof i=="function"?i(c):i;return a.createElement(re,K({key:c,className:j(t,"".concat(t,"-").concat(c),x==null?void 0:x(c)),style:l==null?void 0:l(c),keys:y,motionAppear:!0},v,{onAllRemoved:function(){S(c)}}),function(P,w){var R=P.config,H=P.className,Z=P.style,G=R.key,ee=R.times,ne=R.className,te=R.style;return a.createElement(X,K({},R,{ref:w,prefixCls:t,className:j(H,ne),style:A(A({},Z),te),times:ee,key:G,eventKey:G,onNoticeClose:C}))})})),o)}),he=["getContainer","motion","prefixCls","maxCount","className","style","onAllRemoved"],Ee=function(){return document.body},B=0;function $e(){for(var e={},n=arguments.length,s=new Array(n),t=0;t<n;t++)s[t]=arguments[t];return s.forEach(function(o){o&&Object.keys(o).forEach(function(i){var f=o[i];f!==void 0&&(e[i]=f)})}),e}function Oe(){var e=arguments.length>0&&arguments[0]!==void 0?arguments[0]:{},n=e.getContainer,s=n===void 0?Ee:n,t=e.motion,o=e.prefixCls,i=e.maxCount,f=e.className,x=e.style,l=e.onAllRemoved,u=ie(e,he),g=a.useState(),m=_(g,2),d=m[0],p=m[1],C=a.useRef(),b=a.createElement(be,{container:d,ref:C,prefixCls:o,motion:t,maxCount:i,className:f,style:x,onAllRemoved:l}),N=a.useState([]),$=_(N,2),h=$[0],S=$[1],O=a.useMemo(function(){return{open:function(c){var r=$e(u,c);(r.key===null||r.key===void 0)&&(r.key="rc-notification-".concat(B),B+=1),S(function(y){return[].concat(M(y),[{type:"open",config:r}])})},close:function(c){S(function(r){return[].concat(M(r),[{type:"close",key:c}])})},destroy:function(){S(function(c){return[].concat(M(c),[{type:"destroy"}])})}}},[]);return a.useEffect(function(){p(s())}),a.useEffect(function(){C.current&&h.length&&(h.forEach(function(E){switch(E.type){case"open":C.current.open(E.config);break;case"close":C.current.close(E.key);break;case"destroy":C.current.destroy();break}}),S([]))},[h]),[O,b]}const Pe=e=>{const{componentCls:n,iconCls:s,boxShadow:t,colorText:o,colorSuccess:i,colorError:f,colorWarning:x,colorInfo:l,fontSizeLG:u,motionEaseInOutCirc:g,motionDurationSlow:m,marginXS:d,paddingXS:p,borderRadiusLG:C,zIndexPopup:b,contentPadding:N,contentBg:$}=e,h=`${n}-notice`,S=new z("MessageMoveIn",{"0%":{padding:0,transform:"translateY(-100%)",opacity:0},"100%":{padding:p,transform:"translateY(0)",opacity:1}}),O=new z("MessageMoveOut",{"0%":{maxHeight:e.height,padding:p,opacity:1},"100%":{maxHeight:0,padding:0,opacity:0}}),E={padding:p,textAlign:"center",[`${n}-custom-content > ${s}`]:{verticalAlign:"text-bottom",marginInlineEnd:d,fontSize:u},[`${h}-content`]:{display:"inline-block",padding:N,background:$,borderRadius:C,boxShadow:t,pointerEvents:"all"},[`${n}-success > ${s}`]:{color:i},[`${n}-error > ${s}`]:{color:f},[`${n}-warning > ${s}`]:{color:x},[`${n}-info > ${s},
      ${n}-loading > ${s}`]:{color:l}};return[{[n]:Object.assign(Object.assign({},ue(e)),{color:o,position:"fixed",top:d,width:"100%",pointerEvents:"none",zIndex:b,[`${n}-move-up`]:{animationFillMode:"forwards"},[`
        ${n}-move-up-appear,
        ${n}-move-up-enter
      `]:{animationName:S,animationDuration:m,animationPlayState:"paused",animationTimingFunction:g},[`
        ${n}-move-up-appear${n}-move-up-appear-active,
        ${n}-move-up-enter${n}-move-up-enter-active
      `]:{animationPlayState:"running"},[`${n}-move-up-leave`]:{animationName:O,animationDuration:m,animationPlayState:"paused",animationTimingFunction:g},[`${n}-move-up-leave${n}-move-up-leave-active`]:{animationPlayState:"running"},"&-rtl":{direction:"rtl",span:{direction:"rtl"}}})},{[n]:{[h]:Object.assign({},E)}},{[`${n}-notice-pure-panel`]:Object.assign(Object.assign({},E),{padding:0,textAlign:"start"})}]},Y=ce("Message",e=>{const n=le(e,{height:150});return[Pe(n)]},e=>({zIndexPopup:e.zIndexPopupBase+10,contentBg:e.colorBgElevated,contentPadding:`${(e.controlHeightLG-e.fontSize*e.lineHeight)/2}px ${e.paddingSM}px`}));function Ne(e,n){return{motionName:n??`${e}-move-up`}}function L(e){let n;const s=new Promise(o=>{n=e(()=>{o(!0)})}),t=()=>{n==null||n()};return t.then=(o,i)=>s.then(o,i),t.promise=s,t}var Se=globalThis&&globalThis.__rest||function(e,n){var s={};for(var t in e)Object.prototype.hasOwnProperty.call(e,t)&&n.indexOf(t)<0&&(s[t]=e[t]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var o=0,t=Object.getOwnPropertySymbols(e);o<t.length;o++)n.indexOf(t[o])<0&&Object.prototype.propertyIsEnumerable.call(e,t[o])&&(s[t[o]]=e[t[o]]);return s};const Ie={info:a.createElement(Ce,null),success:a.createElement(ve,null),error:a.createElement(pe,null),warning:a.createElement(ye,null),loading:a.createElement(fe,null)};function q(e){let{prefixCls:n,type:s,icon:t,children:o}=e;return a.createElement("div",{className:j(`${n}-custom-content`,`${n}-${s}`)},t||Ie[s],a.createElement("span",null,o))}function Re(e){const{prefixCls:n,className:s,type:t,icon:o,content:i}=e,f=Se(e,["prefixCls","className","type","icon","content"]),{getPrefixCls:x}=a.useContext(Q),l=n||x("message"),[,u]=Y(l);return a.createElement(X,Object.assign({},f,{prefixCls:l,className:j(s,u,`${l}-notice-pure-panel`),eventKey:"pure",duration:null,content:a.createElement(q,{prefixCls:l,type:t,icon:o},i)}))}var we=globalThis&&globalThis.__rest||function(e,n){var s={};for(var t in e)Object.prototype.hasOwnProperty.call(e,t)&&n.indexOf(t)<0&&(s[t]=e[t]);if(e!=null&&typeof Object.getOwnPropertySymbols=="function")for(var o=0,t=Object.getOwnPropertySymbols(e);o<t.length;o++)n.indexOf(t[o])<0&&Object.prototype.propertyIsEnumerable.call(e,t[o])&&(s[t[o]]=e[t[o]]);return s};const ke=8,je=3,Me=a.forwardRef((e,n)=>{const{top:s,prefixCls:t,getContainer:o,maxCount:i,duration:f=je,rtl:x,transitionName:l,onAllRemoved:u}=e,{getPrefixCls:g,getPopupContainer:m}=a.useContext(Q),d=t||g("message"),[,p]=Y(d),C=()=>({left:"50%",transform:"translateX(-50%)",top:s??ke}),b=()=>j(p,x?`${d}-rtl`:""),N=()=>Ne(d,l),$=a.createElement("span",{className:`${d}-close-x`},a.createElement(de,{className:`${d}-close-icon`})),[h,S]=Oe({prefixCls:d,style:C,className:b,motion:N,closable:!1,closeIcon:$,duration:f,getContainer:()=>(o==null?void 0:o())||(m==null?void 0:m())||document.body,maxCount:i,onAllRemoved:u});return a.useImperativeHandle(n,()=>Object.assign(Object.assign({},h),{prefixCls:d,hashId:p})),S});let U=0;function V(e){const n=a.useRef(null);return[a.useMemo(()=>{const t=l=>{var u;(u=n.current)===null||u===void 0||u.close(l)},o=l=>{if(!n.current){const E=()=>{};return E.then=()=>{},E}const{open:u,prefixCls:g,hashId:m}=n.current,d=`${g}-notice`,{content:p,icon:C,type:b,key:N,className:$,onClose:h}=l,S=we(l,["content","icon","type","key","className","onClose"]);let O=N;return O==null&&(U+=1,O=`antd-message-${U}`),L(E=>(u(Object.assign(Object.assign({},S),{key:O,content:a.createElement(q,{prefixCls:g,type:b,icon:C},p),placement:"top",className:j(b&&`${d}-${b}`,m,$),onClose:()=>{h==null||h(),E()}})),()=>{t(O)}))},f={open:o,destroy:l=>{var u;l!==void 0?t(l):(u=n.current)===null||u===void 0||u.destroy()}};return["info","success","warning","error","loading"].forEach(l=>{const u=(g,m,d)=>{let p;g&&typeof g=="object"&&"content"in g?p=g:p={content:g};let C,b;typeof m=="function"?b=m:(C=m,b=d);const N=Object.assign(Object.assign({onClose:b,duration:C},p),{type:l});return o(N)};f[l]=u}),f},[]),a.createElement(Me,Object.assign({key:"message-holder"},e,{ref:n}))]}function _e(e){return V(e)}let I=null,k=e=>e(),F=[],T={};function Fe(){const{prefixCls:e,getContainer:n,duration:s,rtl:t,maxCount:o,top:i}=T,f=e??W().getPrefixCls("message"),x=(n==null?void 0:n())||document.body;return{prefixCls:f,container:x,duration:s,rtl:t,maxCount:o,top:i}}const Ae=a.forwardRef((e,n)=>{const s=()=>{const{prefixCls:m,container:d,maxCount:p,duration:C,rtl:b,top:N}=Fe();return{prefixCls:m,getContainer:()=>d,maxCount:p,duration:C,rtl:b,top:N}},[t,o]=a.useState(s),[i,f]=V(t),x=W(),l=x.getRootPrefixCls(),u=x.getIconPrefixCls(),g=()=>{o(s)};return a.useEffect(g,[]),a.useImperativeHandle(n,()=>{const m=Object.assign({},i);return Object.keys(m).forEach(d=>{m[d]=function(){return g(),i[d].apply(i,arguments)}}),{instance:m,sync:g}}),a.createElement(ge,{prefixCls:l,iconPrefixCls:u},f)});function D(){if(!I){const e=document.createDocumentFragment(),n={fragment:e};I=n,k(()=>{me(a.createElement(Ae,{ref:s=>{const{instance:t,sync:o}=s||{};Promise.resolve().then(()=>{!n.instance&&t&&(n.instance=t,n.sync=o,D())})}}),e)});return}I.instance&&(F.forEach(e=>{const{type:n,skipped:s}=e;if(!s)switch(n){case"open":{k(()=>{const t=I.instance.open(Object.assign(Object.assign({},T),e.config));t==null||t.then(e.resolve),e.setCloseFn(t)});break}case"destroy":k(()=>{I==null||I.instance.destroy(e.key)});break;default:k(()=>{var t;const o=(t=I.instance)[n].apply(t,M(e.args));o==null||o.then(e.resolve),e.setCloseFn(o)})}}),F=[])}function Te(e){T=Object.assign(Object.assign({},T),e),k(()=>{var n;(n=I==null?void 0:I.sync)===null||n===void 0||n.call(I)})}function De(e){const n=L(s=>{let t;const o={type:"open",config:e,resolve:s,setCloseFn:i=>{t=i}};return F.push(o),()=>{t?k(()=>{t()}):o.skipped=!0}});return D(),n}function He(e,n){const s=L(t=>{let o;const i={type:e,args:n,resolve:t,setCloseFn:f=>{o=f}};return F.push(i),()=>{o?k(()=>{o()}):i.skipped=!0}});return D(),s}function Ke(e){F.push({type:"destroy",key:e}),D()}const Le=["success","info","warning","error","loading"],Ge={open:De,destroy:Ke,config:Te,useMessage:_e,_InternalPanelDoNotUseOrYouWillBeFired:Re},J=Ge;Le.forEach(e=>{J[e]=function(){for(var n=arguments.length,s=new Array(n),t=0;t<n;t++)s[t]=arguments[t];return He(e,s)}});const Qe=J;export{Qe as m};
