(window.webpackJsonp=window.webpackJsonp||[]).push([[0],{114:function(e,t,n){},115:function(e){e.exports={apiKey:"AIzaSyDzHqWsbK8_n21yixU-Eu8HyZsxokY3nmA",authDomain:"dungeon-4a353.firebaseapp.com",databaseURL:"https://dungeon-4a353.firebaseio.com",projectId:"dungeon-4a353",storageBucket:"dungeon-4a353.appspot.com",messagingSenderId:"566690549686",appId:"1:566690549686:web:591525bbe9b8b781"}},116:function(e,t,n){"use strict";n.r(t);var a=n(0),r=n.n(a),c=n(17),l=n.n(c),i=(n(89),n(75)),u=n(14),o=n(7),s=n.n(o),m=n(16),p=n(9),f=n(15),d=n(55),b=n.n(d),v=n(40),h=n(41),E=n(149),g=n(151),j=n(152),O=n(113),y="dungeon-api.itsallonething.com",x=function(){var e=Object(m.a)(s.a.mark(function e(){var t;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("http://".concat(y,"/api"));case 2:return t=e.sent,e.abrupt("return",t.text());case 4:case"end":return e.stop()}},e)}));return function(){return e.apply(this,arguments)}}(),w=function(){var e=Object(m.a)(s.a.mark(function e(){var t,n,a=arguments;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return t=a.length>0&&void 0!==a[0]?a[0]:"",e.next=3,fetch("http://".concat(y,"/api/characters/").concat(t));case 3:return n=e.sent,e.abrupt("return",n.json());case 5:case"end":return e.stop()}},e)}));return function(){return e.apply(this,arguments)}}(),k=function(){var e=Object(m.a)(s.a.mark(function e(t,n){var a;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("http://".concat(y,"/api/characters/").concat(t,"/auth-users?user=").concat(n));case 2:return a=e.sent,e.abrupt("return",a.json());case 4:case"end":return e.stop()}},e)}));return function(t,n){return e.apply(this,arguments)}}(),C={getCharacter:w,saveCharacter:function(){var e=Object(m.a)(s.a.mark(function e(t){return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.prev=0,e.next=3,fetch("http://".concat(y,"/api/characters"),{method:"POST",body:JSON.stringify(t),headers:{"Content-Type":"application/json"}});case 3:e.next=9;break;case 5:return e.prev=5,e.t0=e.catch(0),console.error(e.t0),e.abrupt("return",!1);case 9:return e.abrupt("return",!0);case 10:case"end":return e.stop()}},e,null,[[0,5]])}));return function(t){return e.apply(this,arguments)}}(),watchCharacters:function(e){return new WebSocket("ws://".concat(y,"/api/characters?id=").concat(e.join(),"&watch=true"))},getCharactersByOwner:function(){var e=Object(m.a)(s.a.mark(function e(){var t,n,a=arguments;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return t=a.length>0&&void 0!==a[0]?a[0]:"clayton.yarborough@gmail.com",e.next=3,fetch("http://".concat(y,"/api/characters?owner=").concat(t));case 3:return n=e.sent,e.abrupt("return",n.json());case 5:case"end":return e.stop()}},e)}));return function(){return e.apply(this,arguments)}}(),getWelcome:x,getLevelInfo:function(){var e=Object(m.a)(s.a.mark(function e(t){var n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("http://".concat(y,"/api/level?xp=").concat(t));case 2:return n=e.sent,e.abrupt("return",n.json());case 4:case"end":return e.stop()}},e)}));return function(t){return e.apply(this,arguments)}}(),checkUserAuth:k,getRaces:function(){var e=Object(m.a)(s.a.mark(function e(){var t;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("http://".concat(y,"/api/races"));case 2:return t=e.sent,e.abrupt("return",t.json());case 4:case"end":return e.stop()}},e)}));return function(){return e.apply(this,arguments)}}(),getClasses:function(){var e=Object(m.a)(s.a.mark(function e(t){var n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("http://".concat(y,"/api/classes"));case 2:return n=e.sent,e.abrupt("return",n.json());case 4:case"end":return e.stop()}},e)}));return function(t){return e.apply(this,arguments)}}(),getClass:function(){var e=Object(m.a)(s.a.mark(function e(t){var n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("http://".concat(y,"/api/classes?name=").concat(t));case 2:return n=e.sent,e.abrupt("return",n.json());case 4:case"end":return e.stop()}},e)}));return function(t){return e.apply(this,arguments)}}(),getWeapons:function(){var e=Object(m.a)(s.a.mark(function e(){var t;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("http://".concat(y,"/api/weapons"));case 2:return t=e.sent,e.abrupt("return",t.json());case 4:case"end":return e.stop()}},e)}));return function(){return e.apply(this,arguments)}}(),getSpells:function(){var e=Object(m.a)(s.a.mark(function e(t){var n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("http://".concat(y,"/api/spells?name=").concat(t.join()));case 2:return n=e.sent,e.abrupt("return",n.json());case 4:case"end":return e.stop()}},e)}));return function(t){return e.apply(this,arguments)}}(),getSpellsForLevel:function(){var e=Object(m.a)(s.a.mark(function e(t){var n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,fetch("http://".concat(y,"/api/spells?level=").concat(t));case 2:return n=e.sent,e.abrupt("return",n.json());case 4:case"end":return e.stop()}},e)}));return function(t){return e.apply(this,arguments)}}()};function S(){var e=Object(u.a)(["\n  width: 4.6875em;\n  height: 4.6875em;\n    margin-left: .625em;\n"]);return S=function(){return e},e}function N(){var e=Object(u.a)(["\n  position: fixed;\n  bottom: .313em;\n  right: .313em;\n  display: none;\n  z-index: 101;\n\n  @media screen and (max-width: 62.5em){\n    display: initial;\n  }\n"]);return N=function(){return e},e}function I(){var e=Object(u.a)(["\n  position: fixed;\n  width: 22em;\n  height: 100vh;\n  top: 3.75em;\n  padding: 1em;\n  right: 0em;\n  overflow-y: scroll;\n  z-index: 100;\n\n  @media screen and (max-width: 62.5em){\n    background-color: white;\n    box-shadow: lightgrey 1px 1px .313em;\n    right: -25em;\n    transition: right .3s ease;\n\n    &.open {\n      right: 0em;\n    }\n  }\n"]);return I=function(){return e},e}function L(){var e=Object(u.a)(["\n  flex: 1; \n  display: grid;\n  grid-template-columns: 1fr 25em;\n  margin-bottom: 4.3em;\n  @media screen and (max-width: 62.5em){\n    grid-template-columns: 1fr;\n  }\n"]);return L=function(){return e},e}function W(){var e=Object(u.a)(["\n  display: flex;\n  flex-direction: column;\n"]);return W=function(){return e},e}function P(){var e=Object(u.a)(["\n  display: flex;\n  justify-content: center;\n"]);return P=function(){return e},e}function A(){var e=Object(u.a)(["\n  display: flex;\n  flex-wrap: wrap;\n"]);return A=function(){return e},e}function B(){var e=Object(u.a)(["\n  display: flex;\n  align-items: center;\n  flex-direction: row-reverse;\n"]);return B=function(){return e},e}function D(){var e=Object(u.a)(["\n  display: flex;\n  align-items: center;\n  border-bottom: lightgray solid 1px;\n  margin-bottom: 1em;\n  & h1, h2, h3, h4, p {\n    margin: 0em;\n  }\n  & p {\n    font-size: .8em;\n  }\n"]);return D=function(){return e},e}function z(){var e=Object(u.a)(["\n  position: relative;\n  background: white;\n  box-shadow: lightgray 1px 1px .313em;\n  border-radius: .313em;\n  padding: .625em;\n  display: flex;\n  flex-direction: column;\n  margin: .313em;\n"]);return z=function(){return e},e}function _(){var e=Object(u.a)(["\n  position: fixed;\n  bottom: .313em;\n  right: .313em;\n"]);return _=function(){return e},e}function H(){var e=Object(u.a)(["\n  position: absolute;\n  top: -.8em;\n  right: 0em;\n  z-index: 100;\n"]);return H=function(){return e},e}function F(){var e=Object(u.a)(["\n  flex: 1;\n"]);return F=function(){return e},e}var R=f.a.span(F()),T=f.a.span(H()),U=f.a.span(_()),J=f.a.div(z()),q=f.a.div(D()),G=f.a.span(B()),K=f.a.div(A()),M=f.a.div(P()),V=f.a.div(W()),Y=f.a.div(L()),Z=f.a.div(I()),$=f.a.div(N()),Q=f.a.div(S());var X=function(e){var t=e.user,n=Object(a.useState)("Just a moment..."),c=Object(p.a)(n,2),l=c[0],i=c[1];return Object(a.useEffect)(function(){!function(){var e=Object(m.a)(s.a.mark(function e(){var t;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.prev=0,e.next=3,C.getWelcome();case 3:t=e.sent,i(t),e.next=10;break;case 7:e.prev=7,e.t0=e.catch(0),console.error(e.t0);case 10:case"end":return e.stop()}},e,null,[[0,7]])}));return function(){return e.apply(this,arguments)}}()()},[]),r.a.createElement(E.a,{position:"static"},r.a.createElement(g.a,null,r.a.createElement(v.b,{to:"/"},r.a.createElement("h3",null,l)),r.a.createElement(R,null),t&&r.a.createElement(j.a,null,r.a.createElement(O.a,{alt:t.name,src:t.photo}))))},ee=n(32),te=n(23),ne=n(20),ae=n(156),re=n(162),ce=n(154),le=n(155),ie=n(160),ue=n(166),oe=n(164),se=n(157),me=n(163),pe=n(161);function fe(){var e=Object(u.a)(["\n      display: flex;\n      margin-top: .62em;\n    "]);return fe=function(){return e},e}function de(e){var t=e.character,n=e.save,a=e.add,c=e.linkTo,l=e.open,i=e.highlight;return r.a.createElement(J,{style:i?{background:"lightblue"}:{}},r.a.createElement(q,null,r.a.createElement(V,null,r.a.createElement("h4",null,t.name),r.a.createElement("p",null,t.race," ",t.class)),r.a.createElement(R,null),r.a.createElement(G,null,c&&r.a.createElement(v.b,{to:c},r.a.createElement(j.a,null,r.a.createElement("i",{className:"material-icons"},"fullscreen"))),l&&r.a.createElement(j.a,{onClick:l},r.a.createElement("i",{className:"material-icons"},"fullscreen")),a&&r.a.createElement(j.a,{onClick:a},r.a.createElement("i",{className:"material-icons"},"group_add")),n&&r.a.createElement(j.a,{onClick:n},r.a.createElement("i",{className:"material-icons"},"save")))),r.a.createElement(ve,null,r.a.createElement(re.a,{variant:"outlined",disabled:!0,type:"number",label:"Level",value:t.level}),r.a.createElement(re.a,{variant:"outlined",disabled:!0,type:"number",label:"Hit Points",value:t.hp}),r.a.createElement(re.a,{variant:"outlined",disabled:!0,type:"number",label:"Initiative",value:t.initiative||""})))}var be=de;de.defaultProps={save:null,add:null,linkTo:null,open:null,highlight:!1};var ve=f.a.div(fe());function he(e,t){var n=Math.floor((e-10)/2);return(n=t?n+t:n)<0?"".concat(n):"+".concat(n)}var Ee=[{label:"Acrobatics",check:"dex"},{label:"Animal Handling",check:"wis"},{label:"Arcana",check:"intel"},{label:"Athletics",check:"str"},{label:"Deception",check:"cha"},{label:"History",check:"intel"},{label:"Insight",check:"wis"},{label:"Intimidation",check:"cha"},{label:"Investigation",check:"intel"},{label:"Medicine",check:"wis"},{label:"Nature",check:"intel"},{label:"Perception",check:"wis"},{label:"Performance",check:"cha"},{label:"Persuasion",check:"cha"},{label:"Religion",check:"intel"},{label:"Sleight of Hand",check:"dex"},{label:"Stealth",check:"dex"},{label:"Survival",check:"wis"}];function ge(){var e=Object(u.a)(["\n  display: grid;\n  grid-template-columns: repeat(auto-fill, 150px);\n  grid-gap: 10px;\n  margin: 10px 0px;\n  justify-content: center;\n"]);return ge=function(){return e},e}function je(){var e=Object(u.a)(["\n  display: flex;\n  margin-top: .62em;\n"]);return je=function(){return e},e}function Oe(){var e=Object(u.a)(["\n        position: relative;\n        width: 100vw;\n        display: grid;\n        grid-template-columns: repeat(auto-fill, 22em);\n        grid-gap: .625em;\n        justify-content: center;\n    "]);return Oe=function(){return e},e}var ye=function(e){var t=e.owner,n=Object(a.useState)([]),c=Object(p.a)(n,2),l=c[0],i=c[1],u=Object(a.useState)([]),o=Object(p.a)(u,2),f=o[0],d=o[1],b=Object(a.useState)([]),h=Object(p.a)(b,2),E=h[0],g=h[1],O=Object(a.useState)({}),y=Object(p.a)(O,2),x=y[0],w=y[1],k=Object(a.useState)([]),S=Object(p.a)(k,2),N=S[0],I=S[1],L=Object(a.useState)(!1),W=Object(p.a)(L,2),P=W[0],A=W[1];function B(){return D.apply(this,arguments)}function D(){return(D=Object(m.a)(s.a.mark(function e(){var n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,C.getCharactersByOwner(t);case 2:n=e.sent,i(n||[]);case 4:case"end":return e.stop()}},e)}))).apply(this,arguments)}function z(){return(z=Object(m.a)(s.a.mark(function e(){var t,n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,C.getRaces();case 2:return t=e.sent,g(t),e.next=6,C.getClasses();case 6:n=e.sent,d(n);case 8:case"end":return e.stop()}},e)}))).apply(this,arguments)}function _(e,t){return function(n){return w(Object(te.a)({},x,Object(ee.a)({},e,t?+n.target.value:n.target.value)))}}function H(e){var t=x.proSkills||[],n=t.findIndex(function(t){return t===e.target.value});w(-1===n?Object(te.a)({},x,{proSkills:[].concat(Object(ne.a)(t),[e.target.value])}):Object(te.a)({},x,{proSkills:[].concat(Object(ne.a)(t.slice(0,n)),Object(ne.a)(t.slice(n+1)))}))}function F(){return(F=Object(m.a)(s.a.mark(function e(){var n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return n=Object(te.a)({},x,{owner:t}),e.next=3,C.saveCharacter(n);case 3:w({}),A(!1),B();case 6:case"end":return e.stop()}},e)}))).apply(this,arguments)}Object(a.useEffect)(function(){B()},[t]),Object(a.useEffect)(function(){!P||E.length||f.length||function(){z.apply(this,arguments)}()},[P]);var K=[];return l.forEach(function(e){K.push(r.a.createElement(be,{key:e.id,character:e,highlight:-1!==N.indexOf(e.id),add:function(){return function(e){var t=N.indexOf(e);I(-1!==t?[].concat(Object(ne.a)(N.slice(0,t)),Object(ne.a)(N.slice(t+1))):[].concat(Object(ne.a)(N),[e]))}(e.id)},linkTo:"/character?id=".concat(e.id)}))}),r.a.createElement(xe,null,r.a.createElement(T,null,r.a.createElement(ae.a,{size:"small",color:"secondary",onClick:function(){return A(!P)}},r.a.createElement("i",{className:"material-icons"},P?"close":"add"))),K,P&&r.a.createElement(J,null,r.a.createElement(q,null,r.a.createElement("h4",null,"New Character"),r.a.createElement(R,null),r.a.createElement(G,null,r.a.createElement(j.a,{onClick:function(){return F.apply(this,arguments)}},r.a.createElement("i",{className:"material-icons"},"done")))),r.a.createElement(we,null,r.a.createElement(re.a,{variant:"outlined",label:"Name",onChange:_("name")}),r.a.createElement(Q,null,r.a.createElement(re.a,{variant:"outlined",label:"HP",type:"number",onChange:_("maxHP",!0)}))),r.a.createElement(we,null,r.a.createElement(ce.a,{variant:"outlined",style:{minWidth:120}},r.a.createElement(le.a,{htmlFor:"race"},"Race"),r.a.createElement(ie.a,{value:x.race||"",onChange:function(e){var t=E.find(function(t){return t.name===e.target.value});w(Object(te.a)({},x,{race:t.name,speed:t.speed}))},input:r.a.createElement(ue.a,{id:"race"})},E.map(function(e){return r.a.createElement(oe.a,{key:e.name,value:e.name},e.name)}))),r.a.createElement(R,null),r.a.createElement(ce.a,{variant:"outlined",style:{minWidth:120}},r.a.createElement(le.a,{htmlFor:"class"},"Class"),r.a.createElement(ie.a,{value:x.class||"",onChange:_("class"),input:r.a.createElement(ue.a,{id:"class"})},f.map(function(e){return r.a.createElement(oe.a,{key:e.name,value:e.name},e.name)})))),r.a.createElement(se.a,null),r.a.createElement(q,null,r.a.createElement("p",null,r.a.createElement("strong",null,"Proficient Skills"))),r.a.createElement(ke,null,Ee.map(function(e){var t=e.label;return r.a.createElement(me.a,{key:t,control:r.a.createElement(pe.a,{checked:x.proSkills&&-1!==x.proSkills.indexOf(t)||!1,value:t,onChange:H,color:"primary"}),label:t})}))),!!N.length&&r.a.createElement(U,null,r.a.createElement(v.b,{to:"/character?id=".concat(N.join()),style:{zIndex:10}},r.a.createElement(ae.a,{color:"secondary"},r.a.createElement("i",{className:"material-icons"},"group")))))},xe=f.a.div(Oe()),we=f.a.div(je()),ke=f.a.div(ge()),Ce=n(158),Se=function(e){var t=e.character,n=e.hitDice,a=e.update,c=e.disabled;function l(e,n){return function(r){var c=+r.target.value;c=c>n?n:c,a(Object(te.a)({},t,Object(ee.a)({},e,c)))}}return r.a.createElement(J,null,r.a.createElement(K,{style:{alignItems:"center",justifyContent:"flex-end"}},r.a.createElement(V,null,r.a.createElement("h2",{style:{margin:0}},t.name),r.a.createElement("p",{style:{margin:0}},t.race," ",t.class," (",n,")")),r.a.createElement(R,null),r.a.createElement(Ce.a,{badgeContent:"+".concat(t.proBonus),color:"secondary"},r.a.createElement(Q,null,r.a.createElement(re.a,{variant:"outlined",disabled:!0,type:"number",label:"Level",value:t.level}))),r.a.createElement(Q,null,r.a.createElement(re.a,{variant:"outlined",disabled:c,type:"number",label:"HP/".concat(t.maxHP),value:t.hp,onChange:l("hp")})),r.a.createElement(Q,null,r.a.createElement(re.a,{variant:"outlined",disabled:c,type:"number",label:"AC",value:t.armor,onChange:l("armor")})),r.a.createElement(Q,null,r.a.createElement(re.a,{variant:"outlined",disabled:c,type:"number",label:"Speed",value:t.speed,onChange:l("speed")})),r.a.createElement(Ce.a,{badgeContent:he(t.dex),color:"secondary"},r.a.createElement(Q,null,r.a.createElement(re.a,{variant:"outlined",disabled:c,type:"number",label:"Init",value:t.initiative||"",onChange:l("initiative")})))))};function Ne(){var e=Object(u.a)(["\n  display: grid;\n  grid-template-columns: 1fr 1fr 1fr;\n  grid-template-rows: 1fr 1fr;\n  grid-gap: .625em;\n"]);return Ne=function(){return e},e}var Ie=function(e){var t=e.character,n=e.saves,a=e.update,c=e.disabled;function l(e){return function(n){var r=+n.target.value;a(Object(te.a)({},t,Object(ee.a)({},e,r)))}}return r.a.createElement(J,null,r.a.createElement(q,null,r.a.createElement("h2",null,"Stats")),r.a.createElement(Le,null,r.a.createElement(Ce.a,{badgeContent:he(t.str),color:"secondary"},r.a.createElement(re.a,{variant:"outlined",disabled:c,type:"number",label:"Strength",value:t.str,onChange:l("str")})),r.a.createElement(Ce.a,{badgeContent:he(t.dex),color:"secondary"},r.a.createElement(re.a,{variant:"outlined",disabled:c,type:"number",label:"Dexterity",value:t.dex,onChange:l("dex")})),r.a.createElement(Ce.a,{badgeContent:he(t.con),color:"secondary"},r.a.createElement(re.a,{variant:"outlined",disabled:c,type:"number",label:"Constitution",value:t.con,onChange:l("con")})),r.a.createElement(Ce.a,{badgeContent:he(t.intel),color:"secondary"},r.a.createElement(re.a,{variant:"outlined",disabled:c,type:"number",label:"Intelligence",value:t.intel,onChange:l("intel")})),r.a.createElement(Ce.a,{badgeContent:he(t.wis),color:"secondary"},r.a.createElement(re.a,{variant:"outlined",disabled:c,type:"number",label:"Wisdom",value:t.wis,onChange:l("wis")})),r.a.createElement(Ce.a,{badgeContent:he(t.cha),color:"secondary"},r.a.createElement(re.a,{variant:"outlined",disabled:c,type:"number",label:"Charisma",value:t.cha,onChange:l("cha")}))),r.a.createElement("p",{className:"min-margin"},r.a.createElement("strong",null,"Saving Throws:")," ","+".concat(t.proBonus)," ",n," "))},Le=f.a.div(Ne()),We=function(e){var t=e.character,n=Object(a.useState)(Ee),c=Object(p.a)(n,2),l=c[0],i=c[1],u=Object(a.useState)(""),o=Object(p.a)(u,2),s=o[0],m=o[1];return Object(a.useEffect)(function(){var e=s.toLowerCase(),t=s?Ee.filter(function(t){return 0===t.label.toLowerCase().indexOf(e)}):Ee;i(t)}),r.a.createElement(J,null,r.a.createElement(q,null,r.a.createElement("h2",null,"Skills"),r.a.createElement(R,null),r.a.createElement(G,{style:{marginBottom:"0.3125em"}},r.a.createElement(re.a,{label:"Search",value:s,onChange:function(e){return m(e.target.value)}}))),l.map(function(e){return r.a.createElement("div",{key:e.label},r.a.createElement(K,null,r.a.createElement("p",{style:{margin:"0.3125em"}},e.label),r.a.createElement(R,null),r.a.createElement("h4",{style:{margin:"0.3125em"}},he(t[e.check],function(e,t){return-1!==e.findIndex(function(e){return e===t})}(t.proSkills,e.label)?t.proBonus:0))),r.a.createElement(se.a,null))}))},Pe=function(e){var t=e.proWeapons,n=e.weaponList,c=e.dex,l=e.str,i=e.proBonus,u=e.update,o=e.disabled,f=Object(a.useState)(!1),d=Object(p.a)(f,2),b=d[0],v=d[1],h=Object(a.useState)([]),E=Object(p.a)(h,2),g=E[0],O=E[1],y=Object(a.useState)({}),x=Object(p.a)(y,2),w=x[0],k=x[1];function S(){return(S=Object(m.a)(s.a.mark(function e(){var t;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,C.getWeapons();case 2:t=e.sent,O(t);case 4:case"end":return e.stop()}},e)}))).apply(this,arguments)}return Object(a.useEffect)(function(){b&&!g.length&&function(){S.apply(this,arguments)}()},[b]),r.a.createElement(J,null,r.a.createElement(q,null,r.a.createElement("h2",null,"Weapons"),r.a.createElement(R,null),!o&&r.a.createElement(G,null,r.a.createElement(j.a,{onClick:function(){return v(!b)}},r.a.createElement("i",{className:"material-icons"},b?"close":"add")))),b&&r.a.createElement(K,null,r.a.createElement(ce.a,{variant:"outlined",style:{minWidth:120}},r.a.createElement(le.a,{htmlFor:"class"},"Weapon Select"),r.a.createElement(ie.a,{value:w.name||"",onChange:function(e){var t=g.find(function(t){return t.name===e.target.value});k(t)},input:r.a.createElement(ue.a,{id:"weapon"})},g.map(function(e){return r.a.createElement(oe.a,{key:e.name,value:e.name},e.name)}))),r.a.createElement(R,null),r.a.createElement(j.a,{onClick:function(){u([].concat(Object(ne.a)(n),[w])),v(!1),k({})}},r.a.createElement("i",{className:"material-icons"},"done"))),n.map(function(e){var n=function(e,t){var n=t.split(", ").map(function(e){return e.toLowerCase()});return!!n.find(function(t){return t.indexOf(e.name.toLowerCase())})||!!e.category.split(" ").map(function(e){return e.toLowerCase()}).find(function(e){return!!n.find(function(t){return t.indexOf(e)})})}(e,t)?i:0,a=function(e){return!!e.properties.find(function(e){var t=e.toLowerCase();return-1!==t.indexOf("range")||-1!==t.indexOf("finesse")})}(e)?he(c,n):he(l,n);return r.a.createElement(K,{key:"".concat(e.name)},r.a.createElement(V,null,r.a.createElement("h3",{className:"min-margin"},e.name),r.a.createElement("p",{className:"min-margin"},e.damage_type)),r.a.createElement(R,null),r.a.createElement(Q,null,r.a.createElement(re.a,{variant:"outlined",disabled:!0,label:"Attack",value:a})),r.a.createElement(Q,null,r.a.createElement(re.a,{variant:"outlined",disabled:!0,label:"Damage",value:"".concat(e.damage_dice," ").concat(a)})))}))},Ae=function(e){var t=e.itemList,n=e.gold,c=e.update,l=e.disabled,i=Object(a.useState)(!1),u=Object(p.a)(i,2),o=u[0],s=u[1],m=Object(a.useState)({}),f=Object(p.a)(m,2),d=f[0],b=f[1],v=Object(a.useState)(n),h=Object(p.a)(v,2),E=h[0],g=h[1];function O(e,t){return function(n){b(Object(te.a)({},d,Object(ee.a)({},e,t?+n.target.value:n.target.value)))}}return Object(a.useEffect)(function(){return g(n)},[n]),r.a.createElement(J,null,r.a.createElement(q,null,r.a.createElement("h2",null,"Inventory"),r.a.createElement(R,null),!l&&r.a.createElement(G,null,r.a.createElement(j.a,{onClick:function(){return s(!o)}},r.a.createElement("i",{className:"material-icons"},o?"close":"add")))),o&&r.a.createElement(V,null,r.a.createElement(K,null,r.a.createElement(re.a,{style:{maxWidth:100},variant:"outlined",type:"number",label:"Value (gp)",value:d.goldCost||"",onChange:O("goldCost",!0)}),r.a.createElement(re.a,{style:{maxWidth:150},variant:"outlined",label:"Name",value:d.name||"",onChange:O("name")}),r.a.createElement(R,null),r.a.createElement(j.a,{onClick:function(){c(E,[].concat(Object(ne.a)(t),[d])),s(!1)}},r.a.createElement("i",{className:"material-icons"},"done"))),r.a.createElement(re.a,{variant:"outlined",label:"Description (optional)",value:d.description||"",onChange:O("description")})),r.a.createElement(K,{style:{justifyContent:"flex-end"}},r.a.createElement(re.a,{style:{maxWidth:150},variant:"outlined",type:"number",label:"Gold Pieces",value:E||0,onChange:function(e){var n=+e.target.value;c(n,t),g(n)}})),t.map(function(e,n){return r.a.createElement(V,{key:"".concat(e.name)},r.a.createElement(K,{style:{alignItems:"center"}},r.a.createElement("h4",{className:"min-margin"},e.name),r.a.createElement("p",{className:"min-margin"}," (",e.goldCost,"gp)"),r.a.createElement(R,null),r.a.createElement(j.a,{color:"secondary",onClick:function(){return function(e){c(E,[].concat(Object(ne.a)(t.slice(0,e)),Object(ne.a)(t.slice(e+1))))}(n)}},r.a.createElement("i",{className:"material-icons"},"close"))),e.description&&r.a.createElement("p",{className:"min-margin"},e.description),r.a.createElement(se.a,null))}))},Be=n(168),De=n(159),ze=function(e){var t=e.spell,n=e.close;return r.a.createElement(V,null,r.a.createElement(q,null,r.a.createElement(R,null),r.a.createElement(j.a,{onClick:n},r.a.createElement("i",{className:"material-icons"},"close"))),r.a.createElement(K,null,r.a.createElement(Be.a,{icon:r.a.createElement("i",{className:"material-icons"},"school"),label:t.school,variant:"outlined"}),r.a.createElement(Be.a,{icon:r.a.createElement("i",{className:"material-icons"},"timer"),label:t.casting_time,variant:"outlined"}),r.a.createElement(Be.a,{icon:r.a.createElement("i",{className:"material-icons"},"timelapse"),label:t.duration,variant:"outlined"}),r.a.createElement(Be.a,{icon:r.a.createElement("i",{className:"material-icons"},"wifi_tethering"),label:t.range,variant:"outlined"})),r.a.createElement("p",{className:"min-margin"},t.desc),r.a.createElement(se.a,null),r.a.createElement(K,null,r.a.createElement(R,null)))};function _e(){var e=Object(u.a)(["\n  display: grid;\n  grid-template-columns: repeat(auto-fill, 200px);\n  grid-gap: 10px;\n  margin: 10px 0px;\n  justify-content: center;\n"]);return _e=function(){return e},e}var He=function(e){var t=e.level,n=e.spells,c=e.slots,l=e.addSpell,i=Object(a.useState)([]),u=Object(p.a)(i,2),o=u[0],f=u[1],d=Object(a.useState)([]),b=Object(p.a)(d,2),v=b[0],h=b[1],E=Object(a.useState)([]),g=Object(p.a)(E,2),j=g[0],O=g[1],y=Object(a.useState)(!1),x=Object(p.a)(y,2),w=x[0],k=x[1],S=Object(a.useState)({}),N=Object(p.a)(S,2),I=N[0],L=N[1],W=Object(a.useState)({}),P=Object(p.a)(W,2),A=P[0],B=P[1],D=Object(a.useState)(!1),z=Object(p.a)(D,2),_=z[0],H=z[1],F=Object(a.useState)(!1),T=Object(p.a)(F,2),U=T[0],M=T[1];function Y(){return(Y=Object(m.a)(s.a.mark(function e(){var n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,C.getSpellsForLevel(t);case 2:n=e.sent,O(n.filter(function(e){return!o.find(function(t){return t.slug===e.slug})}));case 4:case"end":return e.stop()}},e)}))).apply(this,arguments)}return Object(a.useEffect)(function(){f(n.filter(function(e){return"Cantrip"===e.level&&"Cantrip"===t||e.level==="".concat(t,"-level")}))},[n]),Object(a.useEffect)(function(){for(var e=[],t=0;t<c;t++)e.push(!1);h(e)},[c]),r.a.createElement(J,null,r.a.createElement(q,null,r.a.createElement("h2",null,t," Level Spells"),r.a.createElement(R,null),r.a.createElement(G,null,v.map(function(e,n){return r.a.createElement(pe.a,{key:"".concat(t,"-slot-").concat(n),checked:e,value:n,onChange:function(){return h([].concat(Object(ne.a)(v.slice(0,n)),[!e],Object(ne.a)(v.slice(n+1))))}})}))),r.a.createElement(V,null,!!A.name&&_&&r.a.createElement(ze,{spell:A,close:function(){B({}),H(!1)}}),r.a.createElement(Fe,null,o.map(function(e){return r.a.createElement(Be.a,{key:e.name,label:e.name,variant:"outlined",onClick:function(){return B(e)},color:e.name===A.name?"primary":"secondary",onDelete:function(){H(!0),B(e)},deleteIcon:r.a.createElement("i",{className:"material-icons"},"info")})})),r.a.createElement(se.a,null)),r.a.createElement(V,null,!!I.name&&U&&r.a.createElement(ze,{spell:I,close:function(){L({}),M(!1)}}),r.a.createElement(Fe,null,w&&j.map(function(e){return r.a.createElement(Be.a,{key:e.name,label:e.name,variant:"outlined",onClick:function(){return L(e)},color:e.name===I.name?"primary":"secondary",onDelete:function(){M(!0),L(e)},deleteIcon:r.a.createElement("i",{className:"material-icons"},"info")})}))),w?r.a.createElement(K,{style:{justifyContent:"center"}},r.a.createElement(De.a,{onClick:function(){k(!1),L({})}},"Cancel"),r.a.createElement(De.a,{disabled:!I.name,variant:"contained",color:"secondary",onClick:function(){return l(I),k(!1),void L({})}},"Learn Spell")):r.a.createElement(De.a,{color:"secondary",onClick:function(){j.length||function(){Y.apply(this,arguments)}(),k(!0)}},"See Spells"))},Fe=f.a.div(_e()),Re=function(e){var t=e.spells,n=e.classInfo,c=e.level,l=e.update,i=Object(a.useState)([]),u=Object(p.a)(i,2),o=u[0],f=u[1];function d(e){l([].concat(Object(ne.a)(t),[e.slug]))}var b=r.a.createElement(He,{level:"Cantrip",slots:0,spells:o||[],addSpell:d}),v=n.Level.map(function(e,t){return t>c-1||!n[e][t]?null:r.a.createElement(He,{key:"".concat(e,"-level-spells"),level:e,slots:+n[e][t],spells:o||[],addSpell:d})});function h(){return(h=Object(m.a)(s.a.mark(function e(t){var n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,C.getSpells(t);case 2:n=e.sent,f(n);case 4:case"end":return e.stop()}},e)}))).apply(this,arguments)}return Object(a.useEffect)(function(){!function(e){h.apply(this,arguments)}(t)},[t]),r.a.createElement(r.a.Fragment,null,b,v)};function Te(){var e=Object(u.a)(["\n    grid-area: etc;\n    display: flex;\n    flex-direction: column;\n  "]);return Te=function(){return e},e}function Ue(){var e=Object(u.a)(["\n    grid-area: eqp;\n  "]);return Ue=function(){return e},e}function Je(){var e=Object(u.a)(["\n    grid-area: wpn;\n  "]);return Je=function(){return e},e}function qe(){var e=Object(u.a)(["\n    grid-area: skill;\n  "]);return qe=function(){return e},e}function Ge(){var e=Object(u.a)(["\n    grid-area: stat;\n  "]);return Ge=function(){return e},e}function Ke(){var e=Object(u.a)(["\n    grid-area: pro;\n  "]);return Ke=function(){return e},e}function Me(){var e=Object(u.a)(['\n    position: relative;\n    display: grid;\n    grid-gap: .625em;\n    grid-template-columns: 18.75em minmax(auto, 15.625em) minmax(auto, 12.5em);\n    grid-template-rows: auto 14.5em auto auto auto;\n    grid-template-areas:\n      "pro pro pro"\n      "skill stat stat"\n      "skill wpn wpn"\n      "skill eqp eqp"\n      "etc etc etc";\n  \n  @media screen and (max-width: 36em){\n        display: flex;\n      flex-direction: column;\n    }\n  ']);return Me=function(){return e},e}var Ve=function(e){var t=e.characterData,n=Object(a.useState)(t),c=Object(p.a)(n,2),l=c[0],i=c[1],u=Object(a.useState)(!1),o=Object(p.a)(u,2),f=o[0],d=o[1],v=Object(a.useState)(!1),h=Object(p.a)(v,2),E=h[0],g=h[1],j=Object(a.useState)(!1),O=Object(p.a)(j,2),y=O[0],x=O[1];function w(e){d(!0),i(e)}function k(){return(k=Object(m.a)(s.a.mark(function e(){return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return d(!1),e.next=3,C.saveCharacter(l);case 3:e.sent||d(!0);case 5:case"end":return e.stop()}},e)}))).apply(this,arguments)}function S(){return(S=Object(m.a)(s.a.mark(function e(t){var n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:if(l){e.next=2;break}return e.abrupt("return");case 2:return e.next=4,C.checkUserAuth(l.id,t.email);case 4:n=e.sent,g(n.authorized);case 6:case"end":return e.stop()}},e)}))).apply(this,arguments)}function N(){return(N=Object(m.a)(s.a.mark(function e(t){var n;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:if(t){e.next=2;break}return e.abrupt("return");case 2:return e.next=4,C.getClass(t);case 4:n=e.sent,x(n);case 6:case"end":return e.stop()}},e)}))).apply(this,arguments)}return Object(a.useEffect)(function(){i(t)},[t]),Object(a.useEffect)(function(){!function(e){S.apply(this,arguments)}(b.a.auth().currentUser),function(e){N.apply(this,arguments)}(l.class)},[l]),r.a.createElement(Ye,null,f&&r.a.createElement(T,null,r.a.createElement(ae.a,{color:"secondary",size:"small",onClick:function(){return k.apply(this,arguments)}},r.a.createElement("i",{className:"material-icons"},"done"))),r.a.createElement(Ze,null,r.a.createElement(Se,{character:l,hitDice:y.hit_dice||"",update:w,disabled:!E})),r.a.createElement($e,null,r.a.createElement(Ie,{character:l,saves:y.prof_saving_throws||"",update:w,disabled:!E})),r.a.createElement(Qe,null,r.a.createElement(We,{character:l})),r.a.createElement(Xe,null,r.a.createElement(Pe,{disabled:!E,proWeapons:y.prof_weapons||"",weaponList:l.weapons||[],dex:l.dex,str:l.str,proBonus:l.proBonus,update:function(e){return w(Object(te.a)({},l,{weapons:e}))}})),r.a.createElement(et,null,r.a.createElement(Ae,{disabled:!E,itemList:l.inventory||[],gold:l.gold,update:function(e,t){return w(Object(te.a)({},l,{gold:e,inventory:t}))}})),r.a.createElement(tt,null,y&&y.spellcasting_ability&&r.a.createElement(Re,{classInfo:y.info,level:l.level,spells:l.spells||[],update:function(e){return w(Object(te.a)({},l,{spells:e}))}})))},Ye=f.a.div(Me()),Ze=f.a.div(Ke()),$e=f.a.div(Ge()),Qe=f.a.div(qe()),Xe=f.a.div(Je()),et=f.a.div(Ue()),tt=f.a.div(Te());var nt=function(e){var t=e.location,n=Object(a.useState)(!1),c=Object(p.a)(n,2),l=c[0],i=c[1],u=Object(a.useState)([]),o=Object(p.a)(u,2),s=o[0],m=o[1],f=Object(a.useState)([]),d=Object(p.a)(f,2),b=d[0],v=d[1],h=Object(a.useState)(null),E=Object(p.a)(h,2),g=E[0],j=E[1];return Object(a.useEffect)(function(){var e=t.search.split("id=")[1].split(",");m(e);var n=C.watchCharacters(e);return n.onmessage=function(e){var t=JSON.parse(e.data).sort(function(e,t){return e.initiative?t.initiative?t.initiative-e.initiative:-1:t.initiative?1:0});v(t)},function(){return n.close()}},[]),Object(a.useEffect)(function(){var e=g&&g.id||s[0];e&&j(function(e){return b.find(function(t){return t.id===e})}(e))},[b,s]),g?r.a.createElement(Y,null,r.a.createElement(M,null,r.a.createElement(Ve,{characterData:g})),b.length>1&&r.a.createElement(r.a.Fragment,null,r.a.createElement(Z,{className:l?"open":""},b.map(function(e){return r.a.createElement(be,{key:e.id,character:e,open:function(){return j(e)},highlight:g.id===e.id})})),r.a.createElement($,null,r.a.createElement(ae.a,{color:"secondary",onClick:function(){return i(!l)}},r.a.createElement("i",{className:"material-icons"},l?"close":"group"))))):null};n(114);function at(){var e=Object(u.a)(["\n  display: flex;\n  margin-top: 1em;\n"]);return at=function(){return e},e}var rt=function(){var e=Object(a.useState)(),t=Object(p.a)(e,2),n=t[0],c=t[1];function l(){return i.apply(this,arguments)}function i(){return(i=Object(m.a)(s.a.mark(function e(){var t,n,a;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return t=new b.a.auth.GoogleAuthProvider,e.next=3,b.a.auth().signInWithPopup(t);case 3:return n=e.sent,a=n.user,e.abrupt("return",{name:a.displayName,email:a.email,photo:a.photoURL});case 6:case"end":return e.stop()}},e)}))).apply(this,arguments)}return Object(a.useEffect)(function(){!function(){var e=Object(m.a)(s.a.mark(function e(){var t;return s.a.wrap(function(e){for(;;)switch(e.prev=e.next){case 0:return e.next=2,l();case 2:t=e.sent,c(t);case 4:case"end":return e.stop()}},e)}));return function(){return e.apply(this,arguments)}}()()},[]),r.a.createElement(v.a,null,r.a.createElement("div",{className:"App "},r.a.createElement(X,{user:n}),n&&r.a.createElement(ct,null,r.a.createElement(h.a,{path:"/",exact:!0,component:function(){return r.a.createElement(ye,{owner:n.email})}}),r.a.createElement(h.a,{path:"/character",exact:!0,component:function(e){return r.a.createElement(nt,e)}}))))},ct=f.a.div(at());Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));var lt=n(115);i.initializeApp(lt),l.a.render(r.a.createElement(rt,null),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then(function(e){e.unregister()})},84:function(e,t,n){e.exports=n(116)},89:function(e,t,n){}},[[84,1,2]]]);
//# sourceMappingURL=main.4bec56dc.chunk.js.map