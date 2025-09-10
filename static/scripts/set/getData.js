fetch('/data/loading', {
    method: 'GET'
}).then(data => data.json())
    .then(data => {
        Main.value = data.mainHeader === "" ? "xx活动" : data.mainHeader;
        Sub.value = data.subHeader === "" ? "主办方：xxx" : data.subHeader;
        Time.value = data.startTime
    })



fetch('/data/pause', {
    method: 'GET'
}).then(data => data.json())
    .then(data => {
            Header.value = data.headerText === "" ? "活动暂停" : data.headerText;
            Describe.value = data.describeText === "" ? "请稍等" : data.describeText;
        }
    )

const inputs = document.querySelectorAll("input");
inputs.forEach((input) => {
    if (input.name !== "time-input") {
        input.addEventListener("mouseenter", function () {
            this.focus();
        });
        input.addEventListener("mouseleave", function () {
            this.blur();
        });
    }
});

