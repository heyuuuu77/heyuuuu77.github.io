console.log('123123123')
const script = document.currentScript;
if (script) {
    const scriptPath = script.src;
    console.log("当前脚本文件的路径是: ", scriptPath);
}