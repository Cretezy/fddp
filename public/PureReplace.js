var prefix = "pr-";
var pageHolderTag = prefix + "data";
var pagesDirectory = "pages";

var currentPage;
var callbacks = [];
/**
 * Start up
 */
function startPureReplace(){
    hashCheck(true); // First load

    window.onhashchange = function () {
        hashCheck(false); // Any changes after
    }
}

function hashCheck(isFirstLoad) {
    // Get hash (page name)
    var hash = window.location.hash.substr(1);
    // Scroll later if first load
    if (isFirstLoad || !scrollToHash(hash)) {
        // Default to "home"
        if (hash == "")
            switchPage('home', function () {
                if (isFirstLoad)
                    scrollToHash(hash);
            });
        else
            switchPage(hash.split("#")[0], function () {
                if (isFirstLoad)
                    scrollToHash(hash);
            });
    }
}

/**
 * Scroll to element based on hash
 * Format: page#element
 *
 * @param hash
 * @returns {boolean} Has scrolled (has an element)
 */
function scrollToHash(hash) {
    if (hash.indexOf("#") > -1) {
        var element = hash.split("#")[1];
        if (element == "top") {
            scrollToTop();
            return false;
        } else {
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
 * @param callback
 * @returns {string} Content
 */
function httpGet(url, callback) {
    var xmlHttp = new XMLHttpRequest();
    xmlHttp.onreadystatechange = function () {
        if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
            callback(xmlHttp.responseText)
        }
    };
    xmlHttp.open("GET", url, true);
    xmlHttp.send();
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
 * @param newPage Page to replace to
 * @param callback
 */
function switchPage(newPage, callback) {
    currentPage = newPage;
    httpGet(pagesDirectory + '/' + newPage + '.html', function (body) {
        document.getElementById(pageHolderTag).innerHTML = body;
        updateLinks();
        window.location.hash = "#" + newPage;
        scrollToTop();
        callbacks.forEach(function(pageLoadCallback){
            pageLoadCallback();
        });
        callback();
    });
}

function onPageLoad(callback){
    callbacks.push(callback)
}