import{r as h,z as j}from"./nox_pyXA7rrAA.js";const m=e=>{let t;const n=new Set,o=(s,p)=>{const c=typeof s=="function"?s(t):s;if(!Object.is(c,t)){const f=t;t=p??typeof c!="object"?c:Object.assign({},t,c),n.forEach(v=>v(t,f))}},r=()=>t,a={setState:o,getState:r,subscribe:s=>(n.add(s),()=>n.delete(s)),destroy:()=>{({BASE_URL:"./",MODE:"production",DEV:!1,PROD:!0,SSR:!1}&&"production")!=="production"&&console.warn("[DEPRECATED] The `destroy` method will be unsupported in a future version. Instead use unsubscribe function returned by subscribe. Everything will be garbage-collected if store is garbage-collected."),n.clear()}};return t=e(o,r,a),a},R=e=>e?m(e):m;var w={exports:{}},D={},b={exports:{}},O={};/**
 * @license React
 * use-sync-external-store-shim.production.min.js
 *
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */var S=h;function V(e,t){return e===t&&(e!==0||1/e===1/t)||e!==e&&t!==t}var $=typeof Object.is=="function"?Object.is:V,_=S.useState,P=S.useEffect,A=S.useLayoutEffect,I=S.useDebugValue;function B(e,t){var n=t(),o=_({inst:{value:n,getSnapshot:t}}),r=o[0].inst,u=o[1];return A(function(){r.value=n,r.getSnapshot=t,y(r)&&u({inst:r})},[e,n,t]),P(function(){return y(r)&&u({inst:r}),e(function(){y(r)&&u({inst:r})})},[e]),I(n),n}function y(e){var t=e.getSnapshot;e=e.value;try{var n=t();return!$(e,n)}catch{return!0}}function C(e,t){return t()}var L=typeof window>"u"||typeof window.document>"u"||typeof window.document.createElement>"u"?C:B;O.useSyncExternalStore=S.useSyncExternalStore!==void 0?S.useSyncExternalStore:L;b.exports=O;var M=b.exports;/**
 * @license React
 * use-sync-external-store-shim/with-selector.production.min.js
 *
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */var E=h,T=M;function k(e,t){return e===t&&(e!==0||1/e===1/t)||e!==e&&t!==t}var q=typeof Object.is=="function"?Object.is:k,z=T.useSyncExternalStore,U=E.useRef,W=E.useEffect,F=E.useMemo,G=E.useDebugValue;D.useSyncExternalStoreWithSelector=function(e,t,n,o,r){var u=U(null);if(u.current===null){var l={hasValue:!1,value:null};u.current=l}else l=u.current;u=F(function(){function s(i){if(!p){if(p=!0,c=i,i=o(i),r!==void 0&&l.hasValue){var d=l.value;if(r(d,i))return f=d}return f=i}if(d=f,q(c,i))return d;var x=o(i);return r!==void 0&&r(d,x)?d:(c=i,f=x)}var p=!1,c,f,v=n===void 0?null:n;return[function(){return s(t())},v===null?void 0:function(){return s(v())}]},[t,n,o,r]);var a=z(e,u[0],u[1]);return W(function(){l.hasValue=!0,l.value=a},[a]),G(a),a};w.exports=D;var H=w.exports;const J=j(H),{useSyncExternalStoreWithSelector:K}=J;function N(e,t=e.getState,n){const o=K(e.subscribe,e.getState,e.getServerState||e.getState,t,n);return h.useDebugValue(o),o}const g=e=>{({BASE_URL:"./",MODE:"production",DEV:!1,PROD:!0,SSR:!1}&&"production")!=="production"&&typeof e!="function"&&console.warn("[DEPRECATED] Passing a vanilla store will be unsupported in a future version. Instead use `import { useStore } from 'zustand'`.");const t=typeof e=="function"?R(e):e,n=(o,r)=>N(t,o,r);return Object.assign(n,t),n},X=e=>e?g(e):g;export{X as c};
