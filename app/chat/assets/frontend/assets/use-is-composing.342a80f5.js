import{r as v,ao as h,w as R,y as p,ap as x,x as E}from"./index.db4b08e0.js";let o=0,a="",f="",d="",c="";const m=v("0px");function H(g){if(typeof document>"u")return;const e=document.documentElement;let n,i=!1;const r=()=>{e.style.marginRight=a,e.style.overflow=f,e.style.overflowX=d,e.style.overflowY=c,m.value="0px"};h(()=>{n=R(g,y=>{if(y){if(!o){const l=window.innerWidth-e.offsetWidth;l>0&&(a=e.style.marginRight,e.style.marginRight=`${l}px`,m.value=`${l}px`),f=e.style.overflow,d=e.style.overflowX,c=e.style.overflowY,e.style.overflow="hidden",e.style.overflowX="hidden",e.style.overflowY="hidden"}i=!0,o++}else o--,o||r(),i=!1},{immediate:!0})}),p(()=>{n==null||n(),i&&(o--,o||r(),i=!1)})}const s=v(!1),w=()=>{s.value=!0},u=()=>{s.value=!1};let t=0;const L=()=>(x&&(E(()=>{t||(window.addEventListener("compositionstart",w),window.addEventListener("compositionend",u)),t++}),p(()=>{t<=1?(window.removeEventListener("compositionstart",w),window.removeEventListener("compositionend",u),t=0):t--})),s);export{L as a,H as u};