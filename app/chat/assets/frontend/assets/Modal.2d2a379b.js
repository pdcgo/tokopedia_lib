import{s as H,r as B,v as te,x as Se,y as we,z as ne,A as se,B as x,C as S,D as w,E as h,G as Pe,H as Ee,I as Be,d as D,J as le,K as M,L as _e,c as F,M as N,O as Ne,P as ce,Q as s,R as O,S as Re,U as de,V as j,W as $e,X as Ae,Y as I,Z as He,k as ve,$ as he,a0 as me,a1 as De,a2 as Ve,a3 as We,a4 as Ke,a5 as Ue,w as pe,a6 as ie,a7 as Te,a8 as be,a9 as A,aa as qe,ab as Xe,ac as Ye,ad as re,ae as Ce,af as Ge,T as Oe,ag as ae,ah as Je,ai as Qe,aj as Ze}from"./index.db4b08e0.js";import{m as Fe,a as eo,d as oo,p as to,F as no,c as io,z as ro,L as ao}from"./index.eb8c481c.js";import{u as so,a as lo}from"./use-is-composing.342a80f5.js";import{i as Ie,h as Me,g as co}from"./utils.5839b03b.js";import{f as uo}from"./fade-in-scale-up.cssr.6fe20355.js";import{e as fo}from"./index.20db2190.js";const L=B(null);function xe(e){if(e.clientX>0||e.clientY>0)L.value={x:e.clientX,y:e.clientY};else{const{target:r}=e;if(r instanceof Element){const{left:t,top:o,width:d,height:u}=r.getBoundingClientRect();t>0||o>0?L.value={x:t+d/2,y:o+u/2}:L.value={x:0,y:0}}else L.value=null}}let E=0,ye=!0;function go(){if(!Ie)return H(B(null));E===0&&te("click",document,xe,!0);const e=()=>{E+=1};return ye&&(ye=Me())?(Se(e),we(()=>{E-=1,E===0&&ne("click",document,xe,!0)})):e(),H(L)}const vo=B(void 0);let _=0;function ke(){vo.value=Date.now()}let ze=!0;function ho(e){if(!Ie)return H(B(!1));const r=B(!1);let t=null;function o(){t!==null&&window.clearTimeout(t)}function d(){o(),r.value=!0,t=window.setTimeout(()=>{r.value=!1},e)}_===0&&te("click",window,ke,!0);const u=()=>{_+=1,te("click",window,d,!0)};return ze&&(ze=Me())?(Se(u),we(()=>{_-=1,_===0&&ne("click",window,ke,!0),ne("click",window,d,!0),o()})):u(),H(r)}const mo={paddingSmall:"12px 16px 12px",paddingMedium:"19px 24px 20px",paddingLarge:"23px 32px 24px",paddingHuge:"27px 40px 28px",titleFontSizeSmall:"16px",titleFontSizeMedium:"18px",titleFontSizeLarge:"18px",titleFontSizeHuge:"18px",closeIconSize:"18px",closeSize:"22px"},po=e=>{const{primaryColor:r,borderRadius:t,lineHeight:o,fontSize:d,cardColor:u,textColor2:c,textColor1:i,dividerColor:g,fontWeightStrong:m,closeIconColor:C,closeIconColorHover:a,closeIconColorPressed:p,closeColorHover:k,closeColorPressed:P,modalColor:y,boxShadow1:f,popoverColor:b,actionColor:v}=e;return Object.assign(Object.assign({},mo),{lineHeight:o,color:u,colorModal:y,colorPopover:b,colorTarget:r,colorEmbedded:v,textColor:c,titleTextColor:i,borderColor:g,actionColor:v,titleFontWeight:m,closeColorHover:k,closeColorPressed:P,closeBorderRadius:t,closeIconColor:C,closeIconColorHover:a,closeIconColorPressed:p,fontSizeSmall:d,fontSizeMedium:d,fontSizeLarge:d,fontSizeHuge:d,boxShadow:f,borderRadius:t})},bo={name:"Card",common:se,self:po},je=bo,Co=x([S("card",`
 font-size: var(--n-font-size);
 line-height: var(--n-line-height);
 display: flex;
 flex-direction: column;
 width: 100%;
 box-sizing: border-box;
 position: relative;
 border-radius: var(--n-border-radius);
 background-color: var(--n-color);
 color: var(--n-text-color);
 word-break: break-word;
 transition: 
 color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 box-shadow .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 `,[w("hoverable",[x("&:hover","box-shadow: var(--n-box-shadow);")]),w("content-segmented",[x(">",[h("content",{paddingTop:"var(--n-padding-bottom)"})])]),w("content-soft-segmented",[x(">",[h("content",`
 margin: 0 var(--n-padding-left);
 padding: var(--n-padding-bottom) 0;
 `)])]),w("footer-segmented",[x(">",[h("footer",{paddingTop:"var(--n-padding-bottom)"})])]),w("footer-soft-segmented",[x(">",[h("footer",`
 padding: var(--n-padding-bottom) 0;
 margin: 0 var(--n-padding-left);
 `)])]),x(">",[S("card-header",`
 box-sizing: border-box;
 display: flex;
 align-items: center;
 font-size: var(--n-title-font-size);
 padding:
 var(--n-padding-top)
 var(--n-padding-left)
 var(--n-padding-bottom)
 var(--n-padding-left);
 `,[h("main",`
 font-weight: var(--n-title-font-weight);
 transition: color .3s var(--n-bezier);
 flex: 1;
 color: var(--n-title-text-color);
 `),h("extra",`
 display: flex;
 align-items: center;
 font-size: var(--n-font-size);
 font-weight: 400;
 transition: color .3s var(--n-bezier);
 color: var(--n-text-color);
 `),h("close",`
 margin: 0 0 0 8px;
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `)]),h("action",`
 box-sizing: border-box;
 transition:
 background-color .3s var(--n-bezier),
 border-color .3s var(--n-bezier);
 background-clip: padding-box;
 background-color: var(--n-action-color);
 `),h("content","flex: 1;"),h("content, footer",`
 box-sizing: border-box;
 padding: 0 var(--n-padding-left) var(--n-padding-bottom) var(--n-padding-left);
 font-size: var(--n-font-size);
 `,[x("&:first-child",{paddingTop:"var(--n-padding-bottom)"})]),h("action",`
 background-color: var(--n-action-color);
 padding: var(--n-padding-bottom) var(--n-padding-left);
 border-bottom-left-radius: var(--n-border-radius);
 border-bottom-right-radius: var(--n-border-radius);
 `)]),S("card-cover",`
 overflow: hidden;
 width: 100%;
 border-radius: var(--n-border-radius) var(--n-border-radius) 0 0;
 `,[x("img",`
 display: block;
 width: 100%;
 `)]),w("bordered",`
 border: 1px solid var(--n-border-color);
 `,[x("&:target","border-color: var(--n-color-target);")]),w("action-segmented",[x(">",[h("action",[x("&:not(:first-child)",{borderTop:"1px solid var(--n-border-color)"})])])]),w("content-segmented, content-soft-segmented",[x(">",[h("content",{transition:"border-color 0.3s var(--n-bezier)"},[x("&:not(:first-child)",{borderTop:"1px solid var(--n-border-color)"})])])]),w("footer-segmented, footer-soft-segmented",[x(">",[h("footer",{transition:"border-color 0.3s var(--n-bezier)"},[x("&:not(:first-child)",{borderTop:"1px solid var(--n-border-color)"})])])])]),Pe(S("card",{background:"var(--n-color-modal)"})),Ee(S("card",{background:"var(--n-color-popover)"})),S("card",[Be({background:"var(--n-color-modal)"})])]),ue={title:String,contentStyle:[Object,String],headerStyle:[Object,String],headerExtraStyle:[Object,String],footerStyle:[Object,String],embedded:Boolean,segmented:{type:[Boolean,Object],default:!1},size:{type:String,default:"medium"},bordered:{type:Boolean,default:!0},closable:{type:Boolean,default:!1},hoverable:Boolean,role:String,onClose:[Function,Array]},xo=de(ue),yo=Object.assign(Object.assign({},M.props),ue),ko=D({name:"Card",props:yo,setup(e){const r=()=>{const{onClose:m}=e;m&&j(m)},{inlineThemeDisabled:t,mergedClsPrefixRef:o,mergedRtlRef:d}=le(e),u=M("Card","-card",Co,je,e,o),c=_e("Card",d,o),i=F(()=>{const{size:m}=e,{self:{color:C,colorModal:a,colorTarget:p,textColor:k,titleTextColor:P,titleFontWeight:y,borderColor:f,actionColor:b,borderRadius:v,lineHeight:R,closeIconColor:z,closeIconColorHover:n,closeIconColorPressed:l,closeColorHover:$,closeColorPressed:T,closeBorderRadius:V,closeIconSize:W,closeSize:K,boxShadow:U,colorPopover:q,colorEmbedded:X,[N("padding",m)]:Y,[N("fontSize",m)]:G,[N("titleFontSize",m)]:J},common:{cubicBezierEaseInOut:Q}}=u.value,{top:Z,left:ee,bottom:oe}=Ne(Y);return{"--n-bezier":Q,"--n-border-radius":v,"--n-color":e.embedded?X:C,"--n-color-modal":a,"--n-color-popover":q,"--n-color-target":p,"--n-text-color":k,"--n-line-height":R,"--n-action-color":b,"--n-title-text-color":P,"--n-title-font-weight":y,"--n-close-icon-color":z,"--n-close-icon-color-hover":n,"--n-close-icon-color-pressed":l,"--n-close-color-hover":$,"--n-close-color-pressed":T,"--n-border-color":f,"--n-box-shadow":U,"--n-padding-top":Z,"--n-padding-bottom":oe,"--n-padding-left":ee,"--n-font-size":G,"--n-title-font-size":J,"--n-close-size":K,"--n-close-icon-size":W,"--n-close-border-radius":V}}),g=t?ce("card",F(()=>e.size[0]),i,e):void 0;return{rtlEnabled:c,mergedClsPrefix:o,mergedTheme:u,handleCloseClick:r,cssVars:t?void 0:i,themeClass:g==null?void 0:g.themeClass,onRender:g==null?void 0:g.onRender}},render(){const{segmented:e,bordered:r,hoverable:t,mergedClsPrefix:o,rtlEnabled:d,onRender:u,$slots:c}=this;return u==null||u(),s("div",{class:[`${o}-card`,this.themeClass,{[`${o}-card--rtl`]:d,[`${o}-card--content${typeof e!="boolean"&&e.content==="soft"?"-soft":""}-segmented`]:e===!0||e!==!1&&e.content,[`${o}-card--footer${typeof e!="boolean"&&e.footer==="soft"?"-soft":""}-segmented`]:e===!0||e!==!1&&e.footer,[`${o}-card--action-segmented`]:e===!0||e!==!1&&e.action,[`${o}-card--bordered`]:r,[`${o}-card--hoverable`]:t}],style:this.cssVars,role:this.role},O(c.cover,i=>i&&s("div",{class:`${o}-card-cover`,role:"none"},i)),O(c.header,i=>i||this.title||this.closable?s("div",{class:`${o}-card-header`,style:this.headerStyle},s("div",{class:`${o}-card-header__main`,role:"heading"},i||[this.title]),O(c["header-extra"],g=>g&&s("div",{class:`${o}-card-header__extra`,style:this.headerExtraStyle},g)),this.closable?s(Re,{clsPrefix:o,class:`${o}-card-header__close`,onClick:this.handleCloseClick,absolute:!0}):null):null),O(c.default,i=>i&&s("div",{class:`${o}-card__content`,style:this.contentStyle,role:"none"},i)),O(c.footer,i=>i&&[s("div",{class:`${o}-card__footer`,style:this.footerStyle,role:"none"},i)]),O(c.action,i=>i&&s("div",{class:`${o}-card__action`,role:"none"},i)))}}),zo={titleFontSize:"18px",padding:"16px 28px 20px 28px",iconSize:"28px",actionSpace:"12px",contentMargin:"8px 0 16px 0",iconMargin:"0 4px 0 0",iconMarginIconTop:"4px 0 8px 0",closeSize:"22px",closeIconSize:"18px",closeMargin:"20px 26px 0 0",closeMarginIconTop:"10px 16px 0 0"},So=e=>{const{textColor1:r,textColor2:t,modalColor:o,closeIconColor:d,closeIconColorHover:u,closeIconColorPressed:c,closeColorHover:i,closeColorPressed:g,infoColor:m,successColor:C,warningColor:a,errorColor:p,primaryColor:k,dividerColor:P,borderRadius:y,fontWeightStrong:f,lineHeight:b,fontSize:v}=e;return Object.assign(Object.assign({},zo),{fontSize:v,lineHeight:b,border:`1px solid ${P}`,titleTextColor:r,textColor:t,color:o,closeColorHover:i,closeColorPressed:g,closeIconColor:d,closeIconColorHover:u,closeIconColorPressed:c,closeBorderRadius:y,iconColor:k,iconColorInfo:m,iconColorSuccess:C,iconColorWarning:a,iconColorError:p,borderRadius:y,titleFontWeight:f})},wo=$e({name:"Dialog",common:se,peers:{Button:Ae},self:So}),Le=wo,fe={icon:Function,type:{type:String,default:"default"},title:[String,Function],closable:{type:Boolean,default:!0},negativeText:String,positiveText:String,positiveButtonProps:Object,negativeButtonProps:Object,content:[String,Function],action:Function,showIcon:{type:Boolean,default:!0},loading:Boolean,bordered:Boolean,iconPlacement:String,onPositiveClick:Function,onNegativeClick:Function,onClose:Function},Po=de(fe),Bo=x([S("dialog",`
 word-break: break-word;
 line-height: var(--n-line-height);
 position: relative;
 background: var(--n-color);
 color: var(--n-text-color);
 box-sizing: border-box;
 margin: auto;
 border-radius: var(--n-border-radius);
 padding: var(--n-padding);
 transition: 
 border-color .3s var(--n-bezier),
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 `,[h("icon",{color:"var(--n-icon-color)"}),w("bordered",{border:"var(--n-border)"}),w("icon-top",[h("close",{margin:"var(--n-close-margin)"}),h("icon",{margin:"var(--n-icon-margin)"}),h("content",{textAlign:"center"}),h("title",{justifyContent:"center"}),h("action",{justifyContent:"center"})]),w("icon-left",[h("icon",{margin:"var(--n-icon-margin)"}),w("closable",[h("title",`
 padding-right: calc(var(--n-close-size) + 6px);
 `)])]),h("close",`
 position: absolute;
 right: 0;
 top: 0;
 margin: var(--n-close-margin);
 transition:
 background-color .3s var(--n-bezier),
 color .3s var(--n-bezier);
 z-index: 1;
 `),h("content",`
 font-size: var(--n-font-size);
 margin: var(--n-content-margin);
 position: relative;
 word-break: break-word;
 `,[w("last","margin-bottom: 0;")]),h("action",`
 display: flex;
 justify-content: flex-end;
 `,[x("> *:not(:last-child)",{marginRight:"var(--n-action-space)"})]),h("icon",{fontSize:"var(--n-icon-size)",transition:"color .3s var(--n-bezier)"}),h("title",`
 transition: color .3s var(--n-bezier);
 display: flex;
 align-items: center;
 font-size: var(--n-title-font-size);
 font-weight: var(--n-title-font-weight);
 color: var(--n-title-text-color);
 `),S("dialog-icon-container",{display:"flex",justifyContent:"center"})]),Pe(S("dialog",`
 width: 446px;
 max-width: calc(100vw - 32px);
 `)),S("dialog",[Be(`
 width: 446px;
 max-width: calc(100vw - 32px);
 `)])]),Ro={default:()=>s(me,null),info:()=>s(me,null),success:()=>s(De,null),warning:()=>s(Ve,null),error:()=>s(We,null)},$o=D({name:"Dialog",alias:["NimbusConfirmCard","Confirm"],props:Object.assign(Object.assign({},M.props),fe),setup(e){const{mergedComponentPropsRef:r,mergedClsPrefixRef:t,inlineThemeDisabled:o}=le(e),d=F(()=>{var a,p;const{iconPlacement:k}=e;return k||((p=(a=r==null?void 0:r.value)===null||a===void 0?void 0:a.Dialog)===null||p===void 0?void 0:p.iconPlacement)||"left"});function u(a){const{onPositiveClick:p}=e;p&&p(a)}function c(a){const{onNegativeClick:p}=e;p&&p(a)}function i(){const{onClose:a}=e;a&&a()}const g=M("Dialog","-dialog",Bo,Le,e,t),m=F(()=>{const{type:a}=e,p=d.value,{common:{cubicBezierEaseInOut:k},self:{fontSize:P,lineHeight:y,border:f,titleTextColor:b,textColor:v,color:R,closeBorderRadius:z,closeColorHover:n,closeColorPressed:l,closeIconColor:$,closeIconColorHover:T,closeIconColorPressed:V,closeIconSize:W,borderRadius:K,titleFontWeight:U,titleFontSize:q,padding:X,iconSize:Y,actionSpace:G,contentMargin:J,closeSize:Q,[p==="top"?"iconMarginIconTop":"iconMargin"]:Z,[p==="top"?"closeMarginIconTop":"closeMargin"]:ee,[N("iconColor",a)]:oe}}=g.value;return{"--n-font-size":P,"--n-icon-color":oe,"--n-bezier":k,"--n-close-margin":ee,"--n-icon-margin":Z,"--n-icon-size":Y,"--n-close-size":Q,"--n-close-icon-size":W,"--n-close-border-radius":z,"--n-close-color-hover":n,"--n-close-color-pressed":l,"--n-close-icon-color":$,"--n-close-icon-color-hover":T,"--n-close-icon-color-pressed":V,"--n-color":R,"--n-text-color":v,"--n-border-radius":K,"--n-padding":X,"--n-line-height":y,"--n-border":f,"--n-content-margin":J,"--n-title-font-size":q,"--n-title-font-weight":U,"--n-title-text-color":b,"--n-action-space":G}}),C=o?ce("dialog",F(()=>`${e.type[0]}${d.value[0]}`),m,e):void 0;return{mergedClsPrefix:t,mergedIconPlacement:d,mergedTheme:g,handlePositiveClick:u,handleNegativeClick:c,handleCloseClick:i,cssVars:o?void 0:m,themeClass:C==null?void 0:C.themeClass,onRender:C==null?void 0:C.onRender}},render(){var e;const{bordered:r,mergedIconPlacement:t,cssVars:o,closable:d,showIcon:u,title:c,content:i,action:g,negativeText:m,positiveText:C,positiveButtonProps:a,negativeButtonProps:p,handlePositiveClick:k,handleNegativeClick:P,mergedTheme:y,loading:f,type:b,mergedClsPrefix:v}=this;(e=this.onRender)===null||e===void 0||e.call(this);const R=u?s(He,{clsPrefix:v,class:`${v}-dialog__icon`},{default:()=>O(this.$slots.icon,n=>n||(this.icon?I(this.icon):Ro[this.type]()))}):null,z=O(this.$slots.action,n=>n||C||m||g?s("div",{class:`${v}-dialog__action`},n||(g?[I(g)]:[this.negativeText&&s(ve,Object.assign({theme:y.peers.Button,themeOverrides:y.peerOverrides.Button,ghost:!0,size:"small",onClick:P},p),{default:()=>I(this.negativeText)}),this.positiveText&&s(ve,Object.assign({theme:y.peers.Button,themeOverrides:y.peerOverrides.Button,size:"small",type:b==="default"?"primary":b,disabled:f,loading:f,onClick:k},a),{default:()=>I(this.positiveText)})])):null);return s("div",{class:[`${v}-dialog`,this.themeClass,this.closable&&`${v}-dialog--closable`,`${v}-dialog--icon-${t}`,r&&`${v}-dialog--bordered`],style:o,role:"dialog"},d?s(Re,{clsPrefix:v,class:`${v}-dialog__close`,onClick:this.handleCloseClick}):null,u&&t==="top"?s("div",{class:`${v}-dialog-icon-container`},R):null,s("div",{class:`${v}-dialog__title`},u&&t==="left"?R:null,he(this.$slots.header,()=>[I(c)])),s("div",{class:[`${v}-dialog__content`,z?"":`${v}-dialog__content--last`]},he(this.$slots.default,()=>[I(i)])),z)}}),To=Ke("n-dialog-provider"),Oo=e=>{const{modalColor:r,textColor2:t,boxShadow3:o}=e;return{color:r,textColor:t,boxShadow:o}},Fo=$e({name:"Modal",common:se,peers:{Scrollbar:Ue,Dialog:Le,Card:je},self:Oo}),Io=Fo,ge=Object.assign(Object.assign({},ue),fe),Mo=de(ge),jo=D({name:"ModalBody",inheritAttrs:!1,props:Object.assign(Object.assign({show:{type:Boolean,required:!0},preset:String,displayDirective:{type:String,required:!0},trapFocus:{type:Boolean,default:!0},autoFocus:{type:Boolean,default:!0},blockScroll:Boolean},ge),{renderMask:Function,onClickoutside:Function,onBeforeLeave:{type:Function,required:!0},onAfterLeave:{type:Function,required:!0},onPositiveClick:{type:Function,required:!0},onNegativeClick:{type:Function,required:!0},onClose:{type:Function,required:!0},onAfterEnter:Function,onEsc:Function}),setup(e){const r=B(null),t=B(null),o=B(e.show),d=B(null),u=B(null);pe(ie(e,"show"),f=>{f&&(o.value=!0)}),so(F(()=>e.blockScroll&&o.value));const c=Te(Fe);function i(){if(c.transformOriginRef.value==="center")return"";const{value:f}=d,{value:b}=u;if(f===null||b===null)return"";if(t.value){const v=t.value.containerScrollTop;return`${f}px ${b+v}px`}return""}function g(f){if(c.transformOriginRef.value==="center")return;const b=c.getMousePosition();if(!b||!t.value)return;const v=t.value.containerScrollTop,{offsetLeft:R,offsetTop:z}=f;if(b){const n=b.y,l=b.x;d.value=-(R-l),u.value=-(z-n-v)}f.style.transformOrigin=i()}function m(f){be(()=>{g(f)})}function C(f){f.style.transformOrigin=i(),e.onBeforeLeave()}function a(){o.value=!1,d.value=null,u.value=null,e.onAfterLeave()}function p(){const{onClose:f}=e;f&&f()}function k(){e.onNegativeClick()}function P(){e.onPositiveClick()}const y=B(null);return pe(y,f=>{f&&be(()=>{const b=f.el;b&&r.value!==b&&(r.value=b)})}),A(eo,r),A(oo,null),A(to,null),{mergedTheme:c.mergedThemeRef,appear:c.appearRef,isMounted:c.isMountedRef,mergedClsPrefix:c.mergedClsPrefixRef,bodyRef:r,scrollbarRef:t,displayed:o,childNodeRef:y,handlePositiveClick:P,handleNegativeClick:k,handleCloseClick:p,handleAfterLeave:a,handleBeforeLeave:C,handleEnter:m}},render(){const{$slots:e,$attrs:r,handleEnter:t,handleAfterLeave:o,handleBeforeLeave:d,preset:u,mergedClsPrefix:c}=this;let i=null;if(!u){if(i=co(e),!i){qe("modal","default slot is empty");return}i=Xe(i),i.props=Ye({class:`${c}-modal`},r,i.props||{})}return this.displayDirective==="show"||this.displayed||this.show?re(s("div",{role:"none",class:`${c}-modal-body-wrapper`},s(Ge,{ref:"scrollbarRef",theme:this.mergedTheme.peers.Scrollbar,themeOverrides:this.mergedTheme.peerOverrides.Scrollbar,contentClass:`${c}-modal-scroll-content`},{default:()=>{var g;return[(g=this.renderMask)===null||g===void 0?void 0:g.call(this),s(no,{disabled:!this.trapFocus,active:this.show,onEsc:this.onEsc,autoFocus:this.autoFocus},{default:()=>{var m;return s(Oe,{name:"fade-in-scale-up-transition",appear:(m=this.appear)!==null&&m!==void 0?m:this.isMounted,onEnter:t,onAfterEnter:this.onAfterEnter,onAfterLeave:o,onBeforeLeave:d},{default:()=>{const C=[[Ce,this.show]],{onClickoutside:a}=this;return a&&C.push([io,this.onClickoutside,void 0,{capture:!0}]),re(this.preset==="confirm"||this.preset==="dialog"?s($o,Object.assign({},this.$attrs,{class:[`${c}-modal`,this.$attrs.class],ref:"bodyRef",theme:this.mergedTheme.peers.Dialog,themeOverrides:this.mergedTheme.peerOverrides.Dialog},ae(this.$props,Po),{"aria-modal":"true"}),e):this.preset==="card"?s(ko,Object.assign({},this.$attrs,{ref:"bodyRef",class:[`${c}-modal`,this.$attrs.class],theme:this.mergedTheme.peers.Card,themeOverrides:this.mergedTheme.peerOverrides.Card},ae(this.$props,xo),{"aria-modal":"true",role:"dialog"}),e):this.childNodeRef=i,C)}})}})]}})),[[Ce,this.displayDirective==="if"||this.displayed||this.show]]):null}}),Lo=x([S("modal-container",`
 position: fixed;
 left: 0;
 top: 0;
 height: 0;
 width: 0;
 display: flex;
 `),S("modal-mask",`
 position: fixed;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 background-color: rgba(0, 0, 0, .4);
 `,[Je({enterDuration:".25s",leaveDuration:".25s",enterCubicBezier:"var(--n-bezier-ease-out)",leaveCubicBezier:"var(--n-bezier-ease-out)"})]),S("modal-body-wrapper",`
 position: fixed;
 left: 0;
 right: 0;
 top: 0;
 bottom: 0;
 overflow: visible;
 `,[S("modal-scroll-content",`
 min-height: 100%;
 display: flex;
 position: relative;
 `)]),S("modal",`
 position: relative;
 align-self: center;
 color: var(--n-text-color);
 margin: auto;
 box-shadow: var(--n-box-shadow);
 `,[uo({duration:".25s",enterScale:".5"})])]),Eo=Object.assign(Object.assign(Object.assign(Object.assign({},M.props),{show:Boolean,unstableShowMask:{type:Boolean,default:!0},maskClosable:{type:Boolean,default:!0},preset:String,to:[String,Object],displayDirective:{type:String,default:"if"},transformOrigin:{type:String,default:"mouse"},zIndex:Number,autoFocus:{type:Boolean,default:!0},trapFocus:{type:Boolean,default:!0},closeOnEsc:{type:Boolean,default:!0},blockScroll:{type:Boolean,default:!0}}),ge),{onEsc:Function,"onUpdate:show":[Function,Array],onUpdateShow:[Function,Array],onAfterEnter:Function,onBeforeLeave:Function,onAfterLeave:Function,onClose:Function,onPositiveClick:Function,onNegativeClick:Function,onMaskClick:Function,internalDialog:Boolean,internalAppear:{type:Boolean,default:void 0},overlayStyle:[String,Object],onBeforeHide:Function,onAfterHide:Function,onHide:Function}),Wo=D({name:"Modal",inheritAttrs:!1,props:Eo,setup(e){const r=B(null),{mergedClsPrefixRef:t,namespaceRef:o,inlineThemeDisabled:d}=le(e),u=M("Modal","-modal",Lo,Io,e,t),c=ho(64),i=go(),g=Qe(),m=e.internalDialog?Te(To,null):null,C=lo();function a(n){const{onUpdateShow:l,"onUpdate:show":$,onHide:T}=e;l&&j(l,n),$&&j($,n),T&&!n&&T(n)}function p(){const{onClose:n}=e;n?Promise.resolve(n()).then(l=>{l!==!1&&a(!1)}):a(!1)}function k(){const{onPositiveClick:n}=e;n?Promise.resolve(n()).then(l=>{l!==!1&&a(!1)}):a(!1)}function P(){const{onNegativeClick:n}=e;n?Promise.resolve(n()).then(l=>{l!==!1&&a(!1)}):a(!1)}function y(){const{onBeforeLeave:n,onBeforeHide:l}=e;n&&j(n),l&&l()}function f(){const{onAfterLeave:n,onAfterHide:l}=e;n&&j(n),l&&l()}function b(n){var l;const{onMaskClick:$}=e;$&&$(n),e.maskClosable&&!((l=r.value)===null||l===void 0)&&l.contains(Ze(n))&&a(!1)}function v(n){var l;(l=e.onEsc)===null||l===void 0||l.call(e),e.show&&e.closeOnEsc&&fo(n)&&!C.value&&a(!1)}A(Fe,{getMousePosition:()=>{if(m){const{clickedRef:n,clickPositionRef:l}=m;if(n.value&&l.value)return l.value}return c.value?i.value:null},mergedClsPrefixRef:t,mergedThemeRef:u,isMountedRef:g,appearRef:ie(e,"internalAppear"),transformOriginRef:ie(e,"transformOrigin")});const R=F(()=>{const{common:{cubicBezierEaseOut:n},self:{boxShadow:l,color:$,textColor:T}}=u.value;return{"--n-bezier-ease-out":n,"--n-box-shadow":l,"--n-color":$,"--n-text-color":T}}),z=d?ce("theme-class",void 0,R,e):void 0;return{mergedClsPrefix:t,namespace:o,isMounted:g,containerRef:r,presetProps:F(()=>ae(e,Mo)),handleEsc:v,handleAfterLeave:f,handleClickoutside:b,handleBeforeLeave:y,doUpdateShow:a,handleNegativeClick:P,handlePositiveClick:k,handleCloseClick:p,cssVars:d?void 0:R,themeClass:z==null?void 0:z.themeClass,onRender:z==null?void 0:z.onRender}},render(){const{mergedClsPrefix:e}=this;return s(ao,{to:this.to,show:this.show},{default:()=>{var r;(r=this.onRender)===null||r===void 0||r.call(this);const{unstableShowMask:t}=this;return re(s("div",{role:"none",ref:"containerRef",class:[`${e}-modal-container`,this.themeClass,this.namespace],style:this.cssVars},s(jo,Object.assign({style:this.overlayStyle},this.$attrs,{ref:"bodyWrapper",displayDirective:this.displayDirective,show:this.show,preset:this.preset,autoFocus:this.autoFocus,trapFocus:this.trapFocus,blockScroll:this.blockScroll},this.presetProps,{onEsc:this.handleEsc,onClose:this.handleCloseClick,onNegativeClick:this.handleNegativeClick,onPositiveClick:this.handlePositiveClick,onBeforeLeave:this.handleBeforeLeave,onAfterEnter:this.onAfterEnter,onAfterLeave:this.handleAfterLeave,onClickoutside:t?void 0:this.handleClickoutside,renderMask:t?()=>{var o;return s(Oe,{name:"fade-in-transition",key:"mask",appear:(o=this.internalAppear)!==null&&o!==void 0?o:this.isMounted},{default:()=>this.show?s("div",{"aria-hidden":!0,ref:"containerRef",class:`${e}-modal-mask`,onClick:this.handleClickoutside}):null})}:void 0}),this.$slots)),[[ro,{zIndex:this.zIndex,enabled:this.show}]])}})}});export{Wo as N,ko as a};
