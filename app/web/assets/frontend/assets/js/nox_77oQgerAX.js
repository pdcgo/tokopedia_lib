import{r as s,C as ue,a1 as pe,a6 as Ce,d as E,e as fe,D as Se,a7 as ge,a8 as he}from"./nox_ophygXhXg.js";import{O as be,a as Oe,u as we,D as xe,b as Ie,c as Pe,T as Ne,g as ve,d as _e}from"./nox_7AoggoAXX.js";import{d as Ee,e as Me}from"./nox_kgkgAgQr7.js";import{F as De}from"./nox_krekpQe7r.js";var ye=globalThis&&globalThis.__rest||function(o,m){var r={};for(var e in o)Object.prototype.hasOwnProperty.call(o,e)&&m.indexOf(e)<0&&(r[e]=o[e]);if(o!=null&&typeof Object.getOwnPropertySymbols=="function")for(var n=0,e=Object.getOwnPropertySymbols(o);n<e.length;n++)m.indexOf(e[n])<0&&Object.prototype.propertyIsEnumerable.call(o,e[n])&&(r[e[n]]=o[e[n]]);return r};const M="SECRET_COMBOBOX_MODE_DO_NOT_USE",Te=(o,m)=>{var r,{prefixCls:e,bordered:n=!0,className:D,rootClassName:x,getPopupContainer:y,popupClassName:T,dropdownClassName:B,listHeight:F=256,placement:d,listItemHeight:W=24,size:z,disabled:C,notFoundContent:I,status:R,showArrow:$,builtinPlacements:j,dropdownMatchSelectWidth:A,popupMatchSelectWidth:f,direction:S}=o,i=ye(o,["prefixCls","bordered","className","rootClassName","getPopupContainer","popupClassName","dropdownClassName","listHeight","placement","listItemHeight","size","disabled","notFoundContent","status","showArrow","builtinPlacements","dropdownMatchSelectWidth","popupMatchSelectWidth","direction"]);const{getPopupContainer:H,getPrefixCls:P,renderEmpty:g,direction:U,virtual:X,popupMatchSelectWidth:k,popupOverflow:G,select:h}=s.useContext(ue),t=P("select",e),L=P(),a=S??U,{compactSize:b,compactItemClassnames:Y}=pe(t,a),[q,N]=we(t),u=s.useMemo(()=>{const{mode:c}=i;if(c!=="combobox")return c===M?"combobox":c},[i.mode]),J=u==="multiple"||u==="tags",v=_e($),K=(r=f??A)!==null&&r!==void 0?r:k,{status:Q,hasFeedback:O,isFormItemInput:V,feedbackIcon:Z}=s.useContext(De),ee=Me(Q,R);let p;I!==void 0?p=I:u==="combobox"?p=null:p=(g==null?void 0:g("Select"))||s.createElement(xe,{componentName:"Select"});const{suffixIcon:te,itemIcon:oe,removeIcon:ne,clearIcon:se}=Ie(Object.assign(Object.assign({},i),{multiple:J,hasFeedback:O,feedbackIcon:Z,showArrow:v,prefixCls:t})),re=Ce(i,["suffixIcon","itemIcon"]),ae=E(T||B,{[`${t}-dropdown-${a}`]:a==="rtl"},x,N),_=fe(c=>{var w;return(w=b??z)!==null&&w!==void 0?w:c}),le=s.useContext(Se),ie=C??le,ce=E({[`${t}-lg`]:_==="large",[`${t}-sm`]:_==="small",[`${t}-rtl`]:a==="rtl",[`${t}-borderless`]:!n,[`${t}-in-form-item`]:V},Ee(t,ee,O),Y,D,x,N),me=s.useMemo(()=>d!==void 0?d:a==="rtl"?"bottomRight":"bottomLeft",[d,a]),de=Pe(j,G);return q(s.createElement(Ne,Object.assign({ref:m,virtual:X,showSearch:h==null?void 0:h.showSearch},re,{dropdownMatchSelectWidth:K,builtinPlacements:de,transitionName:ge(L,he(d),i.transitionName),listHeight:F,listItemHeight:W,mode:u,prefixCls:t,placement:me,direction:a,inputIcon:te,menuItemSelectedIcon:oe,removeIcon:ne,clearIcon:se,notFoundContent:p,className:ce,getPopupContainer:y||H,dropdownClassName:ae,showArrow:O||v,disabled:ie})))},l=s.forwardRef(Te),Be=ve(l);l.SECRET_COMBOBOX_MODE_DO_NOT_USE=M;l.Option=be;l.OptGroup=Oe;l._InternalPanelDoNotUseOrYouWillBeFired=Be;const $e=l;export{$e as S};
