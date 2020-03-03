var plcorpanel = false; // PLC = True, Panel = False
let index = {
    addProtocol: function(name){ 
        let div = document.createElement("div");
        div.className = "Protocol";
        div.innerHTML = `<input type="radio" id="` + name + `" name="protocolchoice" value="` + name + `"><label for="` + name + `">` + name + `</label><br>`
        div.onclick = function(){
            index.setChoices(name);
        };
        //div.onclick = function(){document.getElementById("chosenchoice").innerHTML = name};
        document.getElementById("protocol").appendChild(div);
    },
    about: function(html) {
        let c = document.createElement("div");
        c.innerHTML = html;
        asticode.modaler.setContent(c);
        asticode.modaler.show();
    },
    init: function() {
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            // Listen
            index.listen();

            // Explore default path
            index.explore();

            // Initialise choices
            index.getChoices();
        })
    },
    explore: function(path) {
        // Create message

        let message = {"name": "update"};
        console.log(path)
        if (typeof path !== "undefined") {
            message.payload = ""
        }
        message.payload = path;

        // Send message
        //asticode.loader.show();
        astilectron.sendMessage(message, function(message) {
            // Init
            asticode.loader.hide();

            // Check error
            if (message.name === "error") {
                asticode.notifier.error(message.payload);
                return
            }
            console.log(message.payload)
            if (plcorpanel) {
                document.getElementById("output").innerHTML = message.payload.access;
            } else {
                document.getElementById("output").innerHTML = message.payload.panel; 
            }
            //document.getElementById("output").innerHTML = message.payload.panel;

        })
    },
    listen: function() {
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                    index.about(message.payload);
                    return {payload: "payload"};
                    break;
                case "check.out.menu":
                    asticode.notifier.info(message.payload);
                    break;
                case "plc":
                    plcorpanel = true;
                    index.explore(document.getElementById("data").value);
                    document.getElementById("output").innerHTML = plcorpanel;
                    return {payload: "payload"};
                    break;
                case "panel":
                    plcorpanel = false;
                    index.explore(document.getElementById("data").value);
                    document.getElementById("output").innerHTML = plcorpanel;
                    return {payload: "payload"};
                    break;

            }
        });
    },
    setChoices: function(choice) {
        // Create message
        let message = {"name": "set"};
        message.payload = choice;

        // Send message
        asticode.loader.show();
        astilectron.sendMessage(message, function(message) {
            // Init
            asticode.loader.hide();

        })
        //document.getElementById("chosenchoice").innerHTML=choice;
        
    },
    getChoices: function() {
        // Create message
        let message = {"name": "init"};

        // Send message
        asticode.loader.show();
        astilectron.sendMessage(message, function(message) {
            // Init
            asticode.loader.hide();

            for (let i = 0; i < message.payload.protocol.length; i++) {
                index.addProtocol(message.payload.protocol[i]);
            }

        })
    },
    selectText: function (containerid) {
        if (document.selection) { // IE
            var range = document.body.createTextRange();
            range.moveToElementText(document.getElementById(containerid));
            range.select();
        } else if (window.getSelection) {
            var range = document.createRange();
            range.selectNode(document.getElementById(containerid));
            window.getSelection().removeAllRanges();
            window.getSelection().addRange(range);
            document.execCommand("copy")
        }
    }
};
