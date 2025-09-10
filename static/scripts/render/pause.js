const pauseTitleElement = document.querySelector('#pauseTitle');
const pauseHtmlElement = document.querySelector('#pauseHtml');


fetch('/data/pause')
    .then(res => res.json())
.then(data => {
    pauseTitleElement.innerHTML = data.headerText === "" ? "活动暂停" : data.headerText;
    pauseHtmlElement.innerHTML = data.describeText === "" ? "请稍等" : data.describeText;
})