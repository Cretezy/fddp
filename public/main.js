$(document).ready(function () {
    PureReplace.addPageLoadCallback("import", function () {
        $("#convertForm").submit(function (e) {
            e.preventDefault();
            var form = $(this);
            var formData = new FormData();
            formData.append('messages', form.find('[name="messages"]')[0].files[0]);
            $.ajax({
                url: form.attr("action"),
                data: formData,
                cache: false,
                contentType: false,
                processData: false,
                type: form.attr("method"),
                success: function (data) {
                    PureReplace.switchPage("show", null, data);
                }
            });
        });

        $("#loadForm").submit(function (e) {
            e.preventDefault();
            if (!window.File || !window.FileReader || !window.FileList || !window.Blob) {
                alert('The File APIs are not fully supported in this browser.');
                return;
            }
            var form = $(this);
            var reader = new FileReader();
            reader.onload = function () {
                PureReplace.switchPage("show", null, JSON.parse(window.atob(reader.result.split(",")[1])));
            };
            reader.readAsDataURL(form.find('[name="messages"]')[0].files[0]);
        });
    });

    PureReplace.addPageLoadCallback("show", function (pageData) {
        if (pageData == null) // If nothing to show, go to home
            return PureReplace.switchPage("home");

        $("#fddp-whoami").html(pageData.whoami);

        var threadHolder = $("#fddp-threads");
        var i = 0;
        pageData.threads.forEach(function (thread) {
            i++;
            var messages = "";
            thread.messages.forEach(function (message) {
                messages += E("li", E("strong", message.sender) + ": " + escapeHTML(message.text));
            });
            var threadElement = threadHolder.append(
                E("li",
                E("h1", thread.persons.join(" & "), {"class": "title-" + i}) +
                E("ul", messages, {"class": "messages-" + i})
                    )
            );
            // TODO: fix bug that need "i"
            var titleElement = threadElement.find(".title-" + i);
            var messagesElement = threadElement.find(".messages-" + i);
            titleElement.click(function () {
                messagesElement.toggle();
            });
            messagesElement.hide();
        });
    });

    var escape = document.createElement('textarea');
    escape.style.display = "none";
    function escapeHTML(html) {
        escape.textContent = html;
        return escape.innerHTML;
    }

    PureReplace.start();
});


