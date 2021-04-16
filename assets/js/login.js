$('#login').on('submit', login);

function login(event) {
  event.preventDefault();

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
      "It doesn't seem right...",
      "Please, check your email and password and try again!",
      'error'
    )
  })
}