var prefix = "pr-";
var pageHolderTag = prefix + "data";
var pagesDirectory = "pages";

var currentPage;
/**
 * Start up
 */
(function () {
    hashCheck(true); // First load

    window.onhashchange = function () {
        hashCheck(false); // Any changes after
    }
})();

function hashCheck(isFirstLoad) {
    // Get hash (page name)
    var hash = window.location.hash.substr(1);
    // Scroll later if first load
    if (isFirstLoad || !scroll(hash)) {
        // Default to "home"
        if (hash == "")
            page('home');
        else {
            page(hash.split("#")[0]);
        }
    }

    if (isFirstLoad) {
        scroll(hash);
    }
}

/**
 * Scroll to element based on hash
 * Format: page#element
 *
 * @param hash
 * @returns {boolean} Has scrolled (has an element)
 */
function scroll(hash) {
    console.log("A" + hash);
    if (hash.indexOf("#") > -1) {
        var element = hash.split("#")[1];
        if (element == "top") {
            scrollToTop();
            return false;
        } else {
            console.log(element);
            document.getElementById(element).scrollIntoView(true);
            window.location.hash = "#" + hash;
            return true
        }
    } else {
        scrollToTop();
    }
    return false
}

function scrollToTop() {
    window.scrollTo(0, 0);
}

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
        // Check if link starts with ## tag
        if (link.href.split("#").length > 2) {
            // Make it link to self-page plus element
            link.href = "#" + window.location.hash.split("#")[1] + "#" + link.href.split("#")[2];
        }
    }
}

/**
 * Replace to page
 *
 * @param page Page to replace to
 */
function page(newPage) {
    currentPage = newPage;
    document.getElementById(pageHolderTag).innerHTML = httpGet(pagesDirectory + '/' + newPage + '.html');
    updateLinks();
    window.location.hash = "#" + newPage;
    scrollToTop()
}
