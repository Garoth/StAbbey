goog.provide('st.dom');

/* Holds some convenience functions for manipulating DOM */
goog.scope(function() {
    /* Returns the element with the given ID */
    st.dom.getById = function(id) {
        return document.getElementById(id);
    };

    /* Creates an element */
    st.dom.createElement = function(type, id, classname) {
        var element = document.createElement(type);

        if (id != null) {
            element.id = id;
        }

        if (classname != null) {
            element.className = classname;
        }

        return element;
    };

    /* Creates a div */
    st.dom.createDiv = function(id, classname) {
        return st.dom.createElement('div', id, classname);
    };

    /* Creates an image */
    st.dom.createImg = function(id, classname, src) {
        var img = st.dom.createElement("img", id, classname);
        img.src = src;
        return img;
    };

    /* Removes all of the child elements of an element. */
    st.dom.removeChildren = function(element) {
        while (element.firstChild) {
            element.removeChild(element.firstChild);
        }
    };

    /* Removes an element from the DOM. */
    st.dom.removeElement = function(element) {
        if (element.parentNode) {
            element.parentNode.removeChild(element);
        }
    };

    /* Adds a classname to an element if not already present. */
    st.dom.addClass = function(element, className) {
        var classNames = element.className.split(' ');
        if (!(className in classNames)) {
            element.className += ' ' + className;
        }
    };

    /* Removes a classname from an element. */
    st.dom.removeClass = function(element, className) {
        var classNames = element.className.split(' ');
        var classIndex = classNames.indexOf(className);

        if (classIndex >= 0) {
            classNames.splice(classIndex, 1);
            element.className = classNames.join(' ');
        }
    };

    /* Adds a class to an element if it's not there or removes it if it is. */
    st.dom.toggleClass = function(element, className) {
       var classNames = element.className.split(' ');
       var classIndex = classNames.indexOf(className);

       if (classIndex >= 0) {
          classNames.splice(classIndex, 1);
       } else {
          classNames.push(className);
       }

       element.className = classNames.join(' ');
    };
});
