import{p as i,N as n,g as a}from"./Popover.9be5e582.js";import{W as c,A as l,bx as m,d as h,K as r,r as d,c as v,Q as g}from"./index.db4b08e0.js";const b={padding:"8px 14px"},u=e=>{const{borderRadius:o,boxShadow2:s,baseColor:t}=e;return Object.assign(Object.assign({},b),{borderRadius:o,boxShadow:s,color:m(t,"rgba(0, 0, 0, .85)"),textColor:t})},O=c({name:"Tooltip",common:l,peers:{Popover:i},self:u}),f=O,x=Object.assign(Object.assign({},a),r.props),P=h({name:"Tooltip",props:x,__popover__:!0,setup(e){const o=r("Tooltip","-tooltip",void 0,f,e),s=d(null);return Object.assign(Object.assign({},{syncPosition(){s.value.syncPosition()},setShow(p){s.value.setShow(p)}}),{popoverRef:s,mergedTheme:o,popoverThemeOverrides:v(()=>o.value.self)})},render(){const{mergedTheme:e,internalExtraClass:o}=this;return g(n,Object.assign(Object.assign({},this.$props),{theme:e.peers.Popover,themeOverrides:e.peerOverrides.Popover,builtinThemeOverrides:this.popoverThemeOverrides,internalExtraClass:o.concat("tooltip"),ref:"popoverRef"}),this.$slots)}});export{P as N,f as t};
