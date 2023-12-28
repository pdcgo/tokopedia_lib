import{r as h,I as R}from"./nox_ohyeteyoA.js";const x=t=>{let e;const n=new Set,o=(s,p)=>{const i=typeof s=="function"?s(e):s;if(!Object.is(i,e)){const f=e;e=p??typeof i!="object"?i:Object.assign({},e,i),n.forEach(E=>E(e,f))}},r=()=>e,l={setState:o,getState:r,subscribe:s=>(n.add(s),()=>n.delete(s)),destroy:()=>{n.clear()}};return e=t(o,r,l),l},j=t=>t?x(t):x;var b={exports:{}},g={},w={exports:{}},O={};/**
 * @license React
 * use-sync-external-store-shim.production.min.js
 *
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */var S=h;function A(t,e){return t===e&&(t!==0||1/t===1/e)||t!==t&&e!==e}var V=typeof Object.is=="function"?Object.is:A,q=S.useState,F=S.useEffect,P=S.useLayoutEffect,W=S.useDebugValue;function _(t,e){var n=e(),o=q({inst:{value:n,getSnapshot:e}}),r=o[0].inst,u=o[1];return P(function(){r.value=n,r.getSnapshot=e,y(r)&&u({inst:r})},[t,n,e]),F(function(){return y(r)&&u({inst:r}),t(function(){y(r)&&u({inst:r})})},[t]),W(n),n}function y(t){var e=t.getSnapshot;t=t.value;try{var n=e();return!V(t,n)}catch{return!0}}function $(t,e){return e()}var I=typeof window>"u"||typeof window.document>"u"||typeof window.document.createElement>"u"?$:_;O.useSyncExternalStore=S.useSyncExternalStore!==void 0?S.useSyncExternalStore:I;w.exports=O;var T=w.exports;/**
 * @license React
 * use-sync-external-store-shim/with-selector.production.min.js
 *
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */var v=h,B=T;function C(t,e){return t===e&&(t!==0||1/t===1/e)||t!==t&&e!==e}var L=typeof Object.is=="function"?Object.is:C,M=B.useSyncExternalStore,U=v.useRef,z=v.useEffect,k=v.useMemo,G=v.useDebugValue;g.useSyncExternalStoreWithSelector=function(t,e,n,o,r){var u=U(null);if(u.current===null){var a={hasValue:!1,value:null};u.current=a}else a=u.current;u=k(function(){function s(c){if(!p){if(p=!0,i=c,c=o(c),r!==void 0&&a.hasValue){var d=a.value;if(r(d,c))return f=d}return f=c}if(d=f,L(i,c))return d;var m=o(c);return r!==void 0&&r(d,m)?d:(i=c,f=m)}var p=!1,i,f,E=n===void 0?null:n;return[function(){return s(e())},E===null?void 0:function(){return s(E())}]},[e,n,o,r]);var l=M(t,u[0],u[1]);return z(function(){a.hasValue=!0,a.value=l},[l]),G(l),l};b.exports=g;var H=b.exports;const J=R(H),{useSyncExternalStoreWithSelector:K}=J;function N(t,e=t.getState,n){const o=K(t.subscribe,t.getState,t.getServerState||t.getState,e,n);return h.useDebugValue(o),o}const D=t=>{const e=typeof t=="function"?j(t):t,n=(o,r)=>N(e,o,r);return Object.assign(n,e),n},X=t=>t?D(t):D;export{X as c};
