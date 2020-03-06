var plcorpanel = false; // PLC = True, Panel = False, used for deciding which values will be outputted
let index = {
    addProtocol: function(name){ 
        // From the available protocol choices from go, create radiobuttons that will send messages to go on chosen protocol
        let div = document.createElement("div");
        div.className = "Protocol";
        div.innerHTML = `<input type="radio" id="` + name + `" name="protocolchoice" value="` + name + `"><label for="` + name + `">` + name + `</label><br>`
        div.onclick = function(){
            index.setChoices(name);
        };
        document.getElementById("protocol").appendChild(div);
    },
    about: function(html) {
        // Show about page
        let c = document.createElement("div");
        c.innerHTML = html;
        asticode.modaler.setContent(c);
        asticode.modaler.show();
    },
    init: function() {
        // Initialise asticode, also run the listen function and get the available protocol choices from go
        // Init
        asticode.loader.init();
        asticode.modaler.init();
        asticode.notifier.init();

        // Wait for astilectron to be ready
        document.addEventListener('astilectron-ready', function() {
            // Listen
            index.listen();

            // Explore default path
            // index.explore();

            // Initialise choices
            index.getChoices();
        })
    },
    explore: function(path) {
        // On changes in input textbox, send the data in that textbox to so it can send back converted text
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
        // Listen to messages from go
        astilectron.onMessage(function(message) {
            switch (message.name) {
                case "about":
                    // Show about menu
                    index.about(message.payload);
                    return {payload: "payload"};
                    break;
                case "check.out.menu":
                    // Show "checkout" popup, disabled for now
                    asticode.notifier.info(message.payload);
                    break;
                case "plc":
                    // If go sends plc, show the PLC values in output window
                    plcorpanel = true;
                    index.explore(document.getElementById("data").value);
                    // document.getElementById("output").innerHTML = plcorpanel;
                    return {payload: "payload"};
                    break;
                case "panel":
                    // If go sends panel, show the panel values in output window
                    plcorpanel = false;
                    index.explore(document.getElementById("data").value);
                    // document.getElementById("output").innerHTML = plcorpanel;
                    return {payload: "payload"};
                    break;

            }
        });
    },
    setChoices: function(choice) {
        // Set choice from radiobutton, send message to go so it nows what it will send
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
        // Get all the available protocol choices from go
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
        // Select all the text and then copy to clipboard
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
