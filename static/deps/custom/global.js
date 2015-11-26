var Store = {};

Store.save = function(key,value){
    if (window.localStorage) {
        return localStorage.setItem(key, value);   
    } else {
        document.cookie = key+"="+value;    
    }
}

Store.get = function(key){
    if (window.localStorage) {
        return localStorage.getItem(key);
    } else {
        var cookieValue = null;
        if (document.cookie && document.cookie != '') {
            var cookies = document.cookie.split(';');
            for (var i = 0; i < cookies.length; i++) {
                var cookie = jQuery.trim(cookies[i]);
                if (cookie.substring(0, key.length + 1) == (key + '=')) {
                    cookieValue = decodeURIComponent(cookie.substring(key.length + 1));
                    break;
                }
            }
        }
        return cookieValue;   
    }
}

Store.delete = function(key){
    var del_cookie = function(name) {
        var get_cookie_value = function(name) {
            var cookieValue = null;
            if (document.cookie && document.cookie != '') {
                var cookies = document.cookie.split(';');
                for (var i = 0; i < cookies.length; i++) {
                    var cookie = jQuery.trim(cookies[i]);
                    if (cookie.substring(0, name.length + 1) == (name + '=')) {
                        cookieValue = decodeURIComponent(cookie.substring(name.length + 1));
                        break;
                    }
                }
            }
            return cookieValue;
        }

        var exp = new Date();
        exp.setTime(exp.getTime() + (-1 * 24 * 60 * 60 * 1000));
        var cval = get_cookie_value(name);
        document.cookie = name + "=" + cval + "; expires=" + exp.toGMTString();
    }

    if (window.localStorage) {
        localStorage.removeItem(key);
    } else {
        del_cookie(key);
    }
}