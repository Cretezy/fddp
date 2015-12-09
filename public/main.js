$(document).ready(function () {
    PureReplace.addPageLoadCallback("convert", function () {
        $("#convertForm").submit(function (e) {
            e.preventDefault();
            if (!window.File || !window.FileReader || !window.FileList || !window.Blob) {
                alert('The File APIs are not fully supported in this browser.');
                return;
            }
            var form = $(this);
            var reader = new FileReader();
            reader.onload = function () {
                $.ajax({
                    url: form.attr("action"),
                    data: {"messages": window.atob(reader.result.split(",")[1])},
                    cache: false,
                    type: form.attr("method"),
                    success: function (data) {
                        PureReplace.switchPage("show", function () {
                            $("#whoami").html(data.whoami);
                        });
                    }
                });
            };
            reader.readAsDataURL(form.find("#file")[0].files[0]);
        });
    });
    PureReplace.start();
});

