var PureReplace = {
    // // // // // // // // // // // // //
    //  You may modify these options.   //
    // // // // // // // // // // // // //

    pageHolderTag: "pr-data",
    pagesDirectory: "pages",

    // // // // // // // // // // // // // // // // // // // // // // // //
    //  No not edit past this point unless you know what you are doing.  //
    // // // // // // // // // // // // // // // // // // // // // // // //

    /**
     * Variables
     */
    currentPage: "",
    pageCallbacks: [], globalCallbacks: [],


    /**
     * Start up
     */
    start: function () {
        PureReplace.hashCheck(true); // First load

        window.onhashchange = function () {
            PureReplace.hashCheck(false); // Any changes after
        }
    },

    /**
     * Runs on each changes of hash -> changes page
     *
     * @param isFirstLoad First time loading
     */
    hashCheck: function (isFirstLoad) {
        // Get hash (page name)
        var hash = window.location.hash.substr(1);
        // Scroll later if first load
        // Default to "home"
        if (hash == "") {

            PureReplace.switchPage('home', null);
        } else {
            PureReplace.switchPage(hash.split("#")[0], function () {
                PureReplace.scrollToHash(hash);
            });
        }

    },

    /**
     * Scroll to element based on hash
     * Format: page#element
     *
     * @param hash
     * @returns {boolean} Has scrolled (has an element)
     */
    scrollToHash: function (hash) {
        if (hash.indexOf("#") > -1) { // Has a sub-hash (scroll-to)
            var element = hash.split("#")[1]; // Get sub-hash
            if (element == "top") { // Scroll to top if "top"
                PureReplace.scrollToTop();
                return false;
            } else { // Scroll to element
                document.getElementById(element).scrollIntoView(true);
                window.location.hash = "#" + hash;
                return true;
            }
        } else {
            PureReplace.scrollToTop();
            return false;
        }
    },

    /**
     * Scroll to top of the page
     */
    scrollToTop: function () {
        window.scrollTo(0, 0);
    },

    /**
     * Get content of url
     *
     * @param url URL
     * @param callback
     * @returns {string} Content
     */
    httpGet: function (url, callback) {
        var xmlHttp = new XMLHttpRequest();
        xmlHttp.onreadystatechange = function () {
            if (xmlHttp.readyState == 4 && xmlHttp.status == 200) {
                callback(xmlHttp.responseText);
            }
        };
        xmlHttp.open("GET", url, true);
        xmlHttp.send();
    },

    /**
     * Fix links of current page to links to replace
     */
    updateLinks: function () {
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
    },

    /**
     * Replace to page
     *
     * @param newPage New page
     * @param callback Callback
     */
    switchPage: function (newPage, callback) {
        if (PureReplace.currentPage == newPage) { // Don't double switch
            if (callback != null)
                callback(false); // Run callback once done
        } else {
            PureReplace.currentPage = newPage;
            PureReplace.httpGet(PureReplace.pagesDirectory + '/' + PureReplace.currentPage + '.html', function (body) {
                // Replace with new page
                document.getElementById(PureReplace.pageHolderTag).innerHTML = body;
                PureReplace.updateLinks(); // Fix scroll-to links
                window.location.hash = "#" + newPage; // Set new url hash
                PureReplace.scrollToTop(); // Go to top
                PureReplace.globalCallbacks.forEach(function (globalCallback) {
                    globalCallback(); // Run all global page load callbacks
                });
                PureReplace.pageCallbacks.forEach(function (pageCallback) {
                    if (pageCallback.page == PureReplace.currentPage)
                        pageCallback.callback(); // Run page-specific load callback
                });
                if (callback != null)
                    callback(true); // Run callback once done
            });
        }
    },

    /**
     * Add a callback to be ran after all page load
     *
     * @param callback Callback
     */
    addGlobalLoadCallback: function (callback) {
        PureReplace.globalCallbacks.push(callback)
    },

    /**
     * Add a callback to be ran after a specific page load
     *
     * @param page Page
     * @param callback Callback
     */
    addPageLoadCallback: function (page, callback) {
        PureReplace.pageCallbacks.push({page: page, callback: callback})
    }
};