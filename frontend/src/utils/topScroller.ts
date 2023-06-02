/* eslint-disable @typescript-eslint/no-non-null-assertion */
export const scroller = function (smooth = false) {
    const top = document.getElementById("top")!
    top.scrollIntoView({ behavior: smooth ? "smooth" : "auto" })
}