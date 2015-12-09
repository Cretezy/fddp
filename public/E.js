/**
 * Example usage
 * var list = E.ul(
 *     E.li("This is a list") +
 *     E.li("Hello " + E.strong("world") + "!")
 * )
 *
 */
var E = {
    /**
     *
     * @param tag
     * @param inner
     * @param attributes
     * @returns {string}
     */
    element: function (tag, inner, attributes) {
        var meta = "";
        if (attributes != null) {
            for (var attribute in attributes) {
                if (attributes.hasOwnProperty(attribute))
                    meta += " " + attribute + "=\"" + attributes[attribute] + "\"";
            }
        }
        return "<" + tag + meta + ">" + inner + "</" + tag + ">";
    },

    ul: function (inner, attributes) {
        return E.element("ul", inner, attributes)
    },
    li: function (inner, attributes) {
        return E.element("li", inner, attributes)
    },

    body: function (inner, attributes) {
        return E.element("body", inner, attributes)
    },
    div: function (inner, attributes) {
        return E.element("div", inner, attributes)
    },
    span: function (inner, attributes) {
        return E.element("span", inner, attributes)
    },

    p: function (inner, attributes) {
        return E.element("p", inner, attributes)
    },
    code: function (inner, attributes) {
        return E.element("code", inner, attributes)
    },

    h1: function (inner, attributes) {
        return E.element("h1", inner, attributes)
    },
    h2: function (inner, attributes) {
        return E.element("h2", inner, attributes)
    },
    h3: function (inner, attributes) {
        return E.element("h3", inner, attributes)
    },
    h4: function (inner, attributes) {
        return E.element("h4", inner, attributes)
    },
    h5: function (inner, attributes) {
        return E.element("h5", inner, attributes)
    },
    h6: function (inner, attributes) {
        return E.element("h6", inner, attributes)
    },

    /**
     * @deprecated Use {@link E.strong}
     */
    b: function (inner, attributes) {
        return E.element("b", inner, attributes)
    },
    i: function (inner, attributes) {
        return E.element("i", inner, attributes)
    },
    strong: function (inner, attributes) {
        return E.element("strong", inner, attributes)
    },
    /**
     * @deprecated Use CSS (text-decoration: underline)
     */
    u: function (inner, attributes) {
        return E.element("u", inner, attributes)
    }
};