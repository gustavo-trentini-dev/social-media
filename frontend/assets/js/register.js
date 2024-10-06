$('#register').on('submit', createUser);

function createUser(event) {
  event.preventDefault();
  const password = $('#password').val()
  const confirmPassword = $('#confirm-password').val()

  if (password != confirmPassword) {
    Swal.fire("Ops...", "Passwords are different", "error")
    return
  }

  const name = $('#name').val()
  const email = $('#email').val()
  const nick = $('#nick').val()

  $.ajax({
    url: '/users',
    method: "POST",
    data: {
      name,
      email,
      nick,
      password
    }
  }).done(() => {
    Swal.fire("Sucesso", "user created", "success")
    .then(() => {
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
        Swal.fire("Ops...", "Auth error", "error")
      })
    })
  }).fail((err) => {
    Swal.fire("Ops...", "Error on create user", "error")
  });
}