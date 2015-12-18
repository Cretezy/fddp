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


            var persons = [];
            thread.persons.forEach(function (person) {
                if (person != pageData.whoami) persons.push(person);
            });

            var threadElement = threadHolder.append(
                E("li",
                    E("div",
                        (
                            E("span", persons.join(" & "), {style: "font-size: 2em"}) +
                            E("span", " (" + thread.messages.length + " messages)")
                        ),
                        {"class": "fddp-title-" + i}
                    ) +
                    E("ul", "", {class: ["fddp-messages-" + i, "fddp-hidden"]})
                )
            );

            // TODO: fix bug that need "i"
            var titleElement = threadElement.find(".fddp-title-" + i);
            var messagesElement = threadElement.find(".fddp-messages-" + i);
            titleElement.click(function () {

                if (messagesElement.hasClass("fddp-hidden")) {
                    messagesElement.removeClass("fddp-hidden");
                    if (messagesElement.html() == "") {
                        var messages = "";
                        thread.messages.forEach(function (message) {
                            messages += E("li", E("strong", message.sender, {
                                    class: "tooltip",
                                    time: message.time
                                }) + ": " + escapeHTML(message.text));
                        });
                        messagesElement.html(messages);
                        $('.tooltip').tooltipster({functionInit: function(){
                           return moment($(this).attr("time")).format('MMMM Do YYYY, h:mm a')
                        }});
                    }
                } else {
                    messagesElement.addClass("fddp-hidden");
                }
            });
        });

        var save = $("#fddp-save");
        save.html("Click to save");
        save.click(function () {
            download("messages.json", JSON.stringify(pageData));
        })
    });

    var escape = document.createElement('textarea');
    escape.style.display = "none";
    function escapeHTML(html) {
        escape.textContent = html;
        return escape.innerHTML;
    }

    PureReplace.start();

    var amountScrolled = 300;
    $('body').prepend('<a href="#" class="back-to-top">Back to Top</a>');
    $(window).scroll(function () {
        if ($(window).scrollTop() > amountScrolled) {
            $('a.back-to-top').fadeIn('slow');
        } else {
            $('a.back-to-top').fadeOut('slow');
        }
    });
    $('a.back-to-top').click(function () {
        $('html, body').animate({
            scrollTop: 0
        }, 700);
        return false;
    });
});

// http://stackoverflow.com/a/18197341
function download(filename, text) {
    var element = document.createElement('a');
    element.style.display = 'none';

    element.setAttribute('href', 'data:application/json;charset=utf-8,' + encodeURIComponent(text));
    element.setAttribute('download', filename);

    document.body.appendChild(element);

    element.click();

    document.body.removeChild(element);
}
