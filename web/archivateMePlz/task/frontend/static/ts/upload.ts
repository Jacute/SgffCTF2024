const form: HTMLFormElement | null = document.getElementById("file-upload-form") as HTMLFormElement;
const fileInput: HTMLInputElement | null = document.getElementById("file-upload") as HTMLInputElement;
const output: HTMLElement | null = document.getElementById("output");

form?.addEventListener("submit", async (event) => {
    event.preventDefault();

    if (fileInput?.files?.length === 0) {
        return;
    }

    const formData: FormData = new FormData(form);

    try {
        const response: Response = await fetch("/api/archivator", {
            method: "POST",
            body: formData
        });

        const file:String = (await response.json())["file"];
        const a: HTMLAnchorElement = document.createElement("a");
        a.href = "/api/archivator?file="+file;
        a.textContent = "Download file";
        if (output) {
            output.innerHTML = "";
            output.appendChild(a);
        }

    } catch (error) {
        console.error("Произошла ошибка:", error);
    }
});
