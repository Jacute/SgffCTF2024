"use strict";
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
const form = document.getElementById("file-upload-form");
const fileInput = document.getElementById("file-upload");
const output = document.getElementById("output");
form === null || form === void 0 ? void 0 : form.addEventListener("submit", (event) => __awaiter(void 0, void 0, void 0, function* () {
    var _a;
    event.preventDefault();
    if (((_a = fileInput === null || fileInput === void 0 ? void 0 : fileInput.files) === null || _a === void 0 ? void 0 : _a.length) === 0) {
        return;
    }
    const formData = new FormData(form);
    try {
        const response = yield fetch("/api/archivator", {
            method: "POST",
            body: formData
        });
        const file = (yield response.json())["file"];
        const a = document.createElement("a");
        a.href = "/api/archivator?file=" + file;
        a.textContent = "Download file";
        if (output) {
            output.innerHTML = "";
            output.appendChild(a);
        }
    }
    catch (error) {
        console.error("Произошла ошибка:", error);
    }
}));
