$('#signup-form').on('submit', createUser);

function createUser(event) {
  event.preventDefault();

  if ($('#password').val() != $('#password-check').val()) {
    Swal.fire(
      'Oops...',
      "Your password doesn't match the confirmation field",
      'error'
    )
    return;
  }

  $.ajax({
    url: "/users",
    method: "POST",
    data: {
      name: $('#name').val(),
      email: $('#email').val(),
      username: $('#username').val(),
      password: $('#password').val()
    }
  }).done(function () {
    Swal.fire(
      'Welcome aboard!',
      "You are now part of the best social media ever made!",
      'success'
    ).then(function () {
      $.ajax({
        url: "/login",
        method: "POST",
        data: {
          email: $('#email').val(),
          password: $('#password').val()
        }
      }).done(function () {
        window.location = "/home";
      }).fail(function () {
        Swal.fire(
          'Oops...',
          "Something went wrong while trying to log you in. Please try again later.",
          'error'
        )
      });
    })
  }).fail(function () {
    Swal.fire(
      'Oops...',
      "Something went wrong. We couldn't sign you up at the moment. Please try again later.",
      'error'
    )
  });
}