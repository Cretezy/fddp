/**
 * Construct an HTML element
 *
 * Example usage
 * var list = E("ul"
 *     E("li", "This is a list" +
 *     E("li", "Hello " + E.strong("world") + "!") +
 *     E("li", "Some special element", {"class": ["btn", "btn-success"], "id": "special"}) +
 *     E("li", "Some other special element", "class='btn btn-danger'")
 * )
 *
 * @param {string} tag - HTML Tag
 * @param {string} inner - Element inner content
 * @param  {(Object|string)} [attributes] - HTML Tag's attributes (class, id, etc)
 *
 * @returns {string}
 */
function E(tag, inner, attributes) {
    var meta = "";
    if (typeof attributes !== 'undefined') {
        if (attributes.constructor === Object) {
            for (var attribute in attributes) {
                if (attributes.hasOwnProperty(attribute)) {
                    var value = attributes[attribute];
                    if (value.constructor === Array) {
                        value = value.join(" ");
                    }
                    meta += " " + (attribute + '="' + value + '"');
                }
            }
        }
    }

    return ("<" + tag + meta + ">") + inner + ("</" + tag + ">");
}