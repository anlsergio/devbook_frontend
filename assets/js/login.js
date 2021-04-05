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
    alert("It doesn't seem right... Check your email or password and try again!")
  })
}