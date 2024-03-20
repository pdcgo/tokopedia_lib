import{ay as ue,d as w,o as p,a as v,j as d,u as n,q as N,az as G,aA as ee,r as C,ao as O,c as R,w as U,g as b,h as y,b as a,av as H,k as M,ak as P,aB as K,p as D,m as T,aC as re,au as A,t as x,l as L,n as B,ax as $,aD as ce,aE as pe,F as se,i as oe,aF as de}from"./index.db4b08e0.js";import{N as te}from"./Input.a8ee46ca.js";import{N as ne}from"./Select.74374ef8.js";import{s as S}from"./socket.2fbb088b.js";import{u as _e}from"./roomSummaries.0d271f77.js";import he from"./StateNoData.76e4ded1.js";import{m as F}from"./mutation.640962a3.js";import{c as fe,u as me}from"./TMessage.30ab47ac.js";import{_ as z}from"./TTooltip.vue_vue_type_style_index_0_lang.561ec907.js";import{M as ve}from"./MiniPin.42e30b56.js";import{N as ie}from"./Tooltip.c677ca6c.js";import{N as V}from"./Space.2e1bd2c6.js";import{N as ye}from"./Modal.2d2a379b.js";import{q as Z}from"./queryParser.785886e3.js";import ge from"./TSpinner.f247bbba.js";import"./use-locale.145fb545.js";import"./use-merged-state.119a85f6.js";import"./fade-in-scale-up.cssr.6fe20355.js";import"./Popover.9be5e582.js";import"./index.eb8c481c.js";import"./use-compitable.ba41904e.js";import"./utils.5839b03b.js";import"./index.20db2190.js";import"./get-slot.65c4337d.js";import"./use-is-composing.342a80f5.js";const J=ue("shopFilterStore",{state(){return{name:"",type:"all"}}}),ke={class:"tpd-shopFilteR"},Se=w({__name:"SectionShopFilter",setup(u){const o=J(),e=[{label:"Semua",value:"all"},{label:"Chat",value:"have_chat"},{label:"Online",value:"online"},{label:"Saldo",value:"saldo"},{label:"Ditandai",value:"pinned"}],_=c=>{c===null?o.type="all":o.type=c};return(c,r)=>(p(),v("div",ke,[d(n(te),{clearable:"",class:"sname-filter",size:"small",placeholder:"Nama toko...",value:n(o).name,"onUpdate:value":r[0]||(r[0]=h=>n(o).name=h)},null,8,["value"]),d(n(ne),{class:"fs-12 sstatus-filter",placeholder:"Status",size:"small","consistent-menu-width":!1,options:e,value:n(o).type,"onUpdate:value":_,clearable:""},null,8,["value"])]))}});const be=N(Se,[["__scopeId","data-v-eb74487f"]]),we=a("span",{style:{color:"#555"},class:"fs-11"},"Tidak ada group yang tersedia",-1),$e=w({__name:"SectionGroupSelect",setup(u){const o=G(),e=ee(),_=fe(),c=C(!1),r=C([]);O(()=>{const t=o.query.group;t&&F("addActiveGroup","groups","PUT",null,null,{onSuccess:()=>_.info("Initialize group",2e3)},k=>k+"/"+t)});const h=t=>{const[k,I,q]=o.path.split("/");e.push({path:`/rooms/${q||"chats"}`,query:{group:t}})},f=R(()=>{const t=o.query.group;return t&&r.value.map(k=>k.value).includes(t)?t:null}),m=()=>{H("getlistgroup","groups","GET",null,null,{onSuccess:t=>{o.query.group||t[0]&&e.push({path:o.path,query:{group:t[0]}}),r.value=t.map(k=>({value:k,label:k}))},onLoading:t=>c.value=t})};return U(()=>o.query.group,t=>{F("addActiveGroup","groups","PUT",null,null,{onSuccess:()=>_.info("Changing group",2e3)},k=>k+"/"+t)}),O(()=>{m()}),(t,k)=>(p(),b(n(ne),{style:{position:"sticky",top:"0","background-color":"#fff","z-index":"99"},size:"small",placeholder:"Pilih group",options:r.value,loading:c.value,onUpdateValue:h,value:n(f)},{empty:y(()=>[we]),_:1},8,["options","loading","value"]))}}),Ce=a("i",{class:"fs-11 bi bi-lock"},null,-1),xe=P(" Pin Code "),qe=w({__name:"MiniPinCode",props:{pin:null},setup(u){const o=u;return(e,_)=>(p(),v("div",null,[d(n(ie),{trigger:"hover",class:"fs-11 py-1"},{trigger:y(()=>[d(n(M),{size:"tiny",type:o.pin?"info":"default",secondary:""},{default:y(()=>[Ce]),_:1},8,["type"])]),default:y(()=>[xe]),_:1})]))}}),ae=u=>(D("data-v-e4e1d312"),u=u(),T(),u),Re={class:"mod_Autocomplete fs-13"},Be=ae(()=>a("h3",null,"Setting PIN",-1)),Ne=ae(()=>a("i",{class:"bi bi-shield-lock",style:{"font-size":"43px"}},null,-1)),Ie=P("Batal"),Le=P("Simpan"),Me=w({__name:"MiniPinCodeModal",props:{open:{type:Boolean},shopid:null,pin:null},emits:["update:open","save"],setup(u,{emit:o}){const e=u,_=C(e.pin),c=async r=>{await K.put("akuns/set_pin?shopid="+e.shopid,{pin:r||""}),o("update:open",!1),o("save",r)};return(r,h)=>(p(),b(n(ye),{show:e.open,"mask-closable":!1,"onUpdate:show":h[3]||(h[3]=f=>o("update:open",f))},{default:y(()=>[a("div",Re,[d(n(V),{vertical:"",size:"large"},{default:y(()=>[Be,d(n(V),{justify:"center"},{default:y(()=>[Ne]),_:1}),d(n(te),{value:_.value,"onUpdate:value":h[0]||(h[0]=f=>_.value=f),type:"text",placeholder:"Masukkan PIN"},null,8,["value"]),d(n(V),null,{default:y(()=>[d(n(M),{style:{width:"134px"},size:"small",class:"fs-13",tertiary:"",onClick:h[1]||(h[1]=f=>o("update:open",!1))},{default:y(()=>[Ie]),_:1}),d(n(M),{style:{width:"134px"},size:"small",class:"fs-13",type:"info",onClick:h[2]||(h[2]=()=>c(_.value))},{default:y(()=>[Le]),_:1})]),_:1})]),_:1})])]),_:1},8,["show"]))}});const Pe=N(Me,[["__scopeId","data-v-e4e1d312"]]),De=a("i",{class:"fs-11 bi bi-download"},null,-1),Te=P(" Withdraw "),ze=w({__name:"MiniWithdraw",setup(u){return(o,e)=>(p(),v("div",null,[d(n(ie),{trigger:"hover",class:"fs-11 py-1"},{trigger:y(()=>[d(n(M),{size:"tiny",type:"primary",secondary:""},{default:y(()=>[De]),_:1})]),default:y(()=>[Te]),_:1})]))}}),j=u=>(D("data-v-b8476515"),u=u(),T(),u),Ue={class:"tpd-shop-upper"},Fe={class:"shopcard-detail"},Ge={class:"fs-13 fw-600"},je={class:"tpd-shop-bottom"},We={style:{"margin-left":"-5px"},class:"fs-11 fn-secondary fw-500"},Ee=j(()=>a("i",{class:"bi bi-arrow-repeat fs-11"},null,-1)),Ae=j(()=>a("span",{class:"fs-11 fw-500"},"Reconnect",-1)),Ve=[Ee,Ae],Oe={class:"is-flex"},Ke={key:0,class:"active-badge"},He=j(()=>a("i",null,null,-1)),Je=j(()=>a("span",{style:{color:"rgb(28, 124, 235)"},class:"fs-11 fn-secondary fw-500"},"Aktif",-1)),Qe=[He,Je],Xe={style:{flex:"1"},class:"is-flex is-justify-content-end tpd-shop-status"},Ye={class:"discussion"},Ze={class:"fn-secondary fw-600 fs-11"},es={key:0,class:"paid-order"},ss={class:"fn-secondary fw-600 fs-11"},os={key:0,class:"unpaid-order"},ts={class:"fn-secondary fw-600 fs-11"},ns={key:0,class:"chats"},is={class:"fn-secondary fw-600 fs-11"},as=w({__name:"SectionShopSingle",props:{urlImage:null,shopName:null,saldo:null,shopid:null,isConnected:{type:Boolean},unreadCount:null,confirmedBuyer:null,unConfirmedBuyer:null,discussionCount:null,type:null,url:null,pinned:{type:Boolean},pin:null},emits:["update:pin","withdraw"],setup(u,{emit:o}){const e=u,_=G(),c=ee(),r=C(e.pinned),h=C(!1),f=s=>{F("accountReconnect","groups/reconnect","PUT",null,null,{},function(l){return l+"/"+s})},m=s=>{F("toggleShopPinned","akuns/toggle_pinned","PUT",null,null,{onSuccess:()=>r.value=!r.value},l=>`${l}/${s}`)},t=s=>Object.entries(s).map((l,i)=>`${l[0]}=${l[1]}`).join("&"),k=R(()=>{if(e.url)return e.url;const{query:s}=_,{group:l}=s;return["chats",t({group:l,shop:e.shopid})].join("?")}),I=R(()=>{if(e.url)return e.url;const{query:s}=_,{group:l}=s;return["discussion",t({group:l,shop:e.shopid})].join("?")}),q=R(()=>e.saldo?e.saldo.toLocaleString("id-ID",{style:"currency",currency:"IDR",maximumFractionDigits:0}):"Saldo Kosong"),W=s=>{s.preventDefault();const{x:l,y:i}=s;let g=[{label:r.value?"Hapus Tanda":`Tandai Toko ( ${e.shopName} )`,onClick:()=>{m(e.shopid)},iconClass:r.value?"bi bi-pin-fill fs-11":"bi bi-pin fs-11"},{label:"Buka Room Chat",onClick:()=>{c.push(k.value)},iconClass:"bi bi-chat-left-text fs-11"},{label:"Buka Diskusi",onClick:()=>{c.push(I.value)},iconClass:"bi bi-chat-square-text-fill fs-11",disable:!e.shopid}];e.type==="list"&&(g=[...g,{label:"Koneksikan Ulang",onClick:()=>{f(e.shopid)},iconClass:"bi bi-arrow-repeat fs-12"}]),ce(g,l,i)};return(s,l)=>{const i=re("router-link");return p(),v("div",{class:B(["tpd-shoplist-single",n(_).query.shop==e.shopid&&"tpd-shoplist-single-acv"]),onContextmenu:W},[d(Pe,{open:h.value,"onUpdate:open":l[0]||(l[0]=g=>h.value=g),shopid:e.shopid,pin:e.pin,onSave:l[1]||(l[1]=g=>o("update:pin",g))},null,8,["open","shopid","pin"]),a("div",Ue,[d(i,{class:"shop-link",to:n(k)},null,8,["to"]),d(ve,{pinned:r.value,class:"pin-icon",onClick:l[2]||(l[2]=g=>m(e.shopid))},null,8,["pinned"]),d(qe,{pin:e.pin,class:"pin-code-icon",onClick:l[3]||(l[3]=()=>h.value=!0)},null,8,["pin"]),d(ze,{pin:e.pin,class:"withdraw-icon",onClick:l[4]||(l[4]=g=>o("withdraw"))},null,8,["pin"]),a("div",{style:A({backgroundImage:`url('${e.urlImage}')`}),class:"shopcard-img"},null,4),a("div",Fe,[a("h4",Ge,x(e.shopName),1),a("span",{style:A(n(q)!=="Saldo Kosong"?"color: #22c55e; font-weight: 700":void 0),class:"fs-11 fn-secondary mb-1 fw-400"},x(n(q)),5)])]),d(n(L),{class:"my-1"}),a("div",je,[e.isConnected?(p(),v("div",{key:0,class:B(e.isConnected?"status-badge-connect":"status-badge-disconnect")},[a("i",{class:B(e.isConnected?"bi bi-check fs-11":"bi bi-shield-exclamation fs-10"),style:A({color:e.isConnected?"rgb(37, 224, 121)":"rgb(255, 66, 66)"})},null,6),a("span",We,x(e.isConnected?"Terhubung":"Terputus"),1)],2)):(p(),v("div",{key:1,onClick:l[5]||(l[5]=g=>f(e.shopid)),class:"reconnect-button"},Ve)),a("div",Oe,[n(_).query.shop==e.shopid?(p(),v("div",Ke,Qe)):$("",!0)]),a("div",Xe,[e.discussionCount?(p(),b(z,{key:0,text:`Diskusi ${e.discussionCount}`},{default:y(()=>[a("div",Ye,[a("span",Ze,x(e.discussionCount),1)])]),_:1},8,["text"])):$("",!0),e.confirmedBuyer?(p(),b(z,{key:1,text:`Pesanan Dibayar ${e.confirmedBuyer||0}`},{default:y(()=>[e.confirmedBuyer?(p(),v("div",es,[a("span",ss,x(e.confirmedBuyer),1)])):$("",!0)]),_:1},8,["text"])):$("",!0),e.unConfirmedBuyer?(p(),b(z,{key:2,text:`Belum Dibayar ${e.unConfirmedBuyer||0}`},{default:y(()=>[e.unConfirmedBuyer?(p(),v("div",os,[a("span",ts,x(e.unConfirmedBuyer),1)])):$("",!0)]),_:1},8,["text"])):$("",!0),e.unreadCount?(p(),b(z,{key:3,text:`Pesan baru ${e.unreadCount||0}`},{default:y(()=>[e.unreadCount?(p(),v("div",ns,[a("span",is,x(e.unreadCount),1)])):$("",!0)]),_:1},8,["text"])):$("",!0)])])],34)}}});const le=N(as,[["__scopeId","data-v-b8476515"]]),ls=u=>(D("data-v-9b341872"),u=u(),T(),u),us={class:"tpd-shoplst-cont"},rs={class:"tpd-grpList"},cs=ls(()=>a("i",{class:"bi bi-download"},null,-1)),ps=P("\xA0 auto withdraw "),ds={class:"tpd-shoplist-grpList mt-1"},_s=w({__name:"SectionShopList",setup(u){const o=G(),e=me(),_=C([]),c=C({}),r=_e(),h=s=>{s&&(console.log(s.shopid," -> connect 'kamu ngapain buka console, balik sana!!!'"),r.updateData(s.shopid,{online:!0}),c.value[s.shopid]=!0)},f=s=>{s&&(console.log(s.shopid," disconnect 'kamu ngapain buka console, balik sana!!!'"),r.updateData(s.shopid,{online:!1}),c.value[s.shopid]=!1)},m=()=>{c.value={}},t=s=>{var g,X,Y;const{unreadsSeller:l}=(Y=(X=(g=s==null?void 0:s.event)==null?void 0:g.data)==null?void 0:X.notifications)==null?void 0:Y.chat,{talk_seller:i}=s==null?void 0:s.event.data.notifications.inbox;r.updateData(s.shopid,{unreadsSeller:s.event.data.notifications.chat.unreadsSeller,...s.event.data.notifications.sellerOrderStatus}),_.value=_.value.map(E=>E.id===+s.shopid?{...E,unreadCount:l,diskusi:i}:E)},k=s=>{H("getlistshop","akuns","GET",null,{group_name:s},{onSuccess:l=>{r.activeShopids=l.map(i=>i.id.toString()),l.forEach(i=>r.updateData(i.id.toString(),{online:i.online,saldo:i.saldo})),_.value=l.map(i=>({url:i.profile_url,username:i.shopname,id:i.id,unreadCount:+i.unread_chat,pinned:i.akun_data.pinned,diskusi:i.diskusi,saldo:i.saldo,pin:i.akun_data.pin})),l.forEach(({id:i,online:g})=>{c.value[i]=g})}})},I=s=>{e[s.event.type](`${s.event.name} - ${s.event.message}`)},q=async(s,l)=>{e.info("withdraw "+l),await K.put("akuns/withdraw/"+s)},W=async()=>{e.info("running auto withdraw"),await K.put("akuns/auto_withdraw")};return U(()=>o.query.group,s=>{s&&k(s)},{immediate:!0}),O(()=>{S.on("connected_event",h),S.on("disconnected_event",f),S.on("disconnect",m),S.on("notification",t),S.on("withdraw",I)}),pe(()=>{S.off("notification",t),S.off("connected_event",h),S.off("disconnected_event",f),S.off("disconnect",m),S.off("withdraw",I)}),(s,l)=>(p(),v("div",us,[a("div",rs,[d($e),d(n(M),{size:"small",type:"success",onClick:W},{default:y(()=>[cs,ps]),_:1}),a("div",ds,[n(o).query.group?(p(!0),v(se,{key:1},oe(_.value,i=>(p(),b(le,{"shop-name":i.username,"url-image":i.url,shopid:i.id,"is-connected":c.value[i.id],"unread-count":i.unreadCount,pinned:i.pinned,key:i.id,"discussion-count":i.diskusi,saldo:i.saldo,pin:i.pin,type:"list","onUpdate:pin":g=>i.pin=g||"",onWithdraw:()=>q(i.id,i.username)},null,8,["shop-name","url-image","shopid","is-connected","unread-count","pinned","discussion-count","saldo","pin","onUpdate:pin","onWithdraw"]))),128)):(p(),b(he,{key:0,text:"Pilih group terlebih dahulu","image-size":"small"}))])])]))}});const hs=N(_s,[["__scopeId","data-v-9b341872"]]),Q=u=>(D("data-v-045cca49"),u=u(),T(),u),fs={class:"section-room-mode"},ms=Q(()=>a("h4",{class:"fs-12 fw-500 fn-secondary"},"Room mode",-1)),vs={class:"section-room-mode-selection"},ys=Q(()=>a("span",{class:"fs-12 fw-500"},"Chat",-1)),gs=[ys],ks=Q(()=>a("span",{class:"fs-12 fw-500"},"Data Order",-1)),Ss=[ks],bs=w({__name:"SectionShopRoomMode",setup(u){const o=G(),e=R(()=>{const{group:c}=o.query,r=Z({group:c});return["/rooms/orders",r].join("?")}),_=R(()=>{const{group:c,shop:r,chat:h}=o.query,f=Z({group:c,shop:r,chat:h});return["/rooms/chats",f].join("?")});return(c,r)=>(p(),v("div",fs,[ms,a("div",vs,[a("button",{onClick:r[0]||(r[0]=h=>c.$router.push(n(_))),type:"button",class:B([n(o).path.includes("chats")&&"active-mode"])},gs,2),a("button",{onClick:r[1]||(r[1]=h=>c.$router.push(n(e))),type:"button",class:B([n(o).path.includes("orders")&&"active-mode"])},Ss,2)])]))}});const ws=N(bs,[["__scopeId","data-v-045cca49"]]),$s=u=>(D("data-v-2ce072f4"),u=u(),T(),u),Cs={key:0,class:"shops_Results scrollbar"},xs={key:1,style:{width:"100%","align-self":"center","text-align":"center"},class:"fs-12 fw-500"},qs=$s(()=>a("span",null,"Hasil tidak ditemukan",-1)),Rs=[qs],Bs={key:2,style:{width:"100%","align-self":"center"}},Ns=w({__name:"SectionFilterResult",props:{show:{type:Boolean}},setup(u){const o=u,e=J(),_=C([]),c=C(!1),h=de(f=>{const m={};switch(f.name&&(m.name=f.name),f.type){case"have_chat":m.have_chat=!0;break;case"online":m.online=!0;break;case"pinned":m.pinned=!0;break;case"saldo":m.saldo=!0;break;default:delete m.have_chat,delete m.online,delete m.pinned,delete m.saldo}H("getlistshop","akuns","GET",null,m,{onSuccess:t=>_.value=t,onLoading:t=>c.value=t})},300);return U(()=>e.name,f=>{f&&(c.value=!0,h(e.$state))}),U(()=>e.type,f=>{c.value=!0,h(e.$state)}),(f,m)=>(p(),v("div",{class:B([o.show&&"shop_ResultWrapper-on ","shop_ResultWrapper"])},[!c.value&&_.value.length?(p(),v("div",Cs,[(p(!0),v(se,null,oe(_.value,t=>(p(),b(le,{key:t.id,"is-connected":t.online,"shop-name":t.shopname,shopid:t.id,"url-image":t.profile_url,saldo:t.saldo,"unread-count":+t.unread_chat,pinned:t.akun_data.pinned,url:`/rooms/chats?group=${t.akun_data.groups[0].name}&shop=${t.id}`},null,8,["is-connected","shop-name","shopid","url-image","saldo","unread-count","pinned","url"]))),128))])):!c.value&&!_.value.length?(p(),v("div",xs,Rs)):(p(),v("div",Bs,[d(ge,{text:"Sedang memuat hasil pencarian"})]))],2))}});const Is=N(Ns,[["__scopeId","data-v-2ce072f4"]]),Ls={class:"tpd-shopSectioN"},Ms={class:"tpd-shopSection scrollbar pr-1"},io=w({__name:"TRoomShopSection",setup(u){const o=J();return(e,_)=>(p(),v("div",Ls,[d(n(L),{class:"my-0 mb-3"}),d(be),n(o).name||n(o).type!=="all"?(p(),b(n(L),{key:0,class:"my-3"})):$("",!0),d(Is,{show:Boolean(n(o).name)||n(o).type!=="all"},null,8,["show"]),d(n(L),{class:"my-3"}),d(ws),d(n(L),{class:"my-3"}),a("div",Ms,[d(hs)])]))}});export{io as default};