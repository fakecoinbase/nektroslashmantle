"use strict";
//
import { create_element, dcTN } from "./../util.js";
import * as api from "./../api/index.js";

//
customElements.define("x-uonline-role", class UOnlineRole extends HTMLElement {
    constructor() {
        super();
    }
    async connectedCallback() {
        this._uid = this.getAttribute("uuid");
        this._pos = this.getAttribute("position");
        const o = await api.M.roles.get(this._uid);
        const n = o.name === undefined ? this.getAttribute("name") : o.name;
        this.appendChild(create_element("div", [["data-count","0"]], [dcTN(n)]));
        this.appendChild(create_element("ul"));
    }
    get count() {
        return parseInt(this.children[0].dataset.count, 10);
    }
    set count(x) {
        this.children[0].dataset.count = x.toString();
    }
    addUser(uid) {
        this.children[1].appendChild(create_element("x-uonline-user", [["uuid",uid],["data-role",this._uid]]));
        this.count += 1;
    }
    removeUser(uid) {
        this.querySelector(`x-uonline-user[uuid="${uid}"]`).remove();
        this.count -= 1;
    }
});