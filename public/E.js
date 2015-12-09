/**
 * Example usage
 * var list = E("ul"
 *     E("li", "This is a list") +
 *     E("li", "Hello " + E.strong("world") + "!")
 * )
 *
 * @param tag
 * @param inner
 * @param [attributes]
 * @returns {string}
 */
function E(tag, inner, attributes) {
    var meta = "";
    if (typeof attributes !== 'undefined') {
        for (var attribute in attributes) {
            if (attributes.hasOwnProperty(attribute))
                meta += " " + (attribute + '="' + attributes[attribute] + '"');
        }
    }

    return ("<" + tag + meta + ">") + inner + ("</" + tag + ">");
}
