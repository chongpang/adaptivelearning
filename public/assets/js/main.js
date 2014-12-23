// Begin of helper functions
Array.prototype.contains = function(obj) {
    var i = this.length;
    while (i--) {
        if (this[i] === obj) {
            return true;
        }
    }
    return false;
}

Array.prototype.remove = function(obj) {
  var l = this.length;
  for(var i = l - 1; i >= 0; i--) {
      if(this[i] === obj) {
         this.splice(i, 1);
      }
  }
}
// End of helper functions

// Handle the change event of the mulitiple selection control
var prerequisites = [];
function ChangeSelection(form, selection) {
  var v = "";
  for (i = 0; i < selection.options.length; i++) {
      if (selection.options[i].selected) {
        if(!prerequisites.contains(selection.options[i].value)){
          prerequisites.push(selection.options[i].value);
        }
      }else{
          prerequisites.remove(selection.options[i].value);
      }
  }
}

// Global event/function call defined here.
$(document).ready(function() {
  // Submit learning object data.

    $('#creatlo-form').submit(function( event ) {
      // Stop form from submitting normally
      event.preventDefault();
     
      // Get some values from elements on the page:
      var $form = $( this );
      var param = {};
       $(form.serializeArray()).each(function(i, v) {
    		param[v.name] = v.value;
    	});

      param['prerequisites'] = prerequisites;

      var json = JSON.stringify(param);
      var url = $form.attr( "action" ); 
      // Send the data using post
      var posting = $.post( url, json); 
      // Put the results in a div
      posting.done(function( data ) {
    	$.alert(data);
      });
    });
});
// End if global event/function call definition.