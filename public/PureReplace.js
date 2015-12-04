var linkTag = "data-page";
var pageHolderTag = "pager";
var pagesDirectory = "pages";

/**
 * Start up
 */
(function () {
    // Go back to page name after "#"
    var hash = window.location.hash.substr(1);
    // Default to "home"
    if (hash == "")
        page('home');
    else
        page(hash);
})();

/**
 * Get content of url
 *
 * @param url URL
 * @returns {string} Content
 */
function httpGet(url) {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.open("GET", url, false);
    xmlHttp.send(null);
    return xmlHttp.responseText;
}

/**
 * Fix links of current page to links to replace
 */
function updateLinks() {
    // Find all links
    var links = document.getElementsByTagName('a');
    for (var i = 0; i < links.length; i++) {
        var link = links[i];
        // Check if link to replace
        if (link.hasAttribute(linkTag)) {
            // Make clickable
            link.href = "#" + link.getAttribute(linkTag);
            link.addEventListener('click', function (e) {
                page(e.target.getAttribute(linkTag));
                e.preventDefault();
            });
        }
    }
}

/**
 * Replace to page
 *
 * @param page Page to replace to
 */
function page(page) {
    document.getElementById(pageHolderTag).innerHTML = httpGet(pagesDirectory + '/' + page + '.html');
    updateLinks();
    window.location.hash = "#" + page;
}

