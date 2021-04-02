$('#signup-form').on('submit', createUser);

function createUser(event) {
  event.preventDefault();

  if ($('#password').val() != $('#password-check').val()) {
    alert("Password doesn't match the confirmation field");
    return;
  }

  $.ajax({
    url: "/users",
    method: "POST",
    data: {
      name: $('#name').val(),
      email: $('#email').val(),
      username: $('#username').val(),
      password: $('#password').val(),
    }
  }).done(function () {
    alert("Welcome aboard " + $('#username') + "!");
  }).fail(function (error) {
    console.log(error);
    alert("Something went wrong. Please try again.");
  });
}