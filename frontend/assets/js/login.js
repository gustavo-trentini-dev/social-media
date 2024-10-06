$('#login').on('submit', login);

function login(event) {
  event.preventDefault();
  const email = $('#email').val()
  const password = $('#password').val()

  $.ajax({
    url: '/login',
    method: "POST",
    data: {
      email,
      password
    }
  }).done(() => {
    window.location = '/home'
  }).fail(() => {
    Swal.fire('Ops...', 'Error doing login', 'error')
  });
}