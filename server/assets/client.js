const deleteHandler = e => {
  $(e.target).parent().attr("data-id");
};

const genPair = pair =>
  `<div data-id=${pair.id} class="container--center g--12">
  <div contenteditable class="card g--2">${pair.name}</div>
  <div contenteditable class="card g--4">${pair.value}</div>
  <div class="icon-holder">
<i class="material-icons md-48">close</i>
</div>

</div>`;

const render = html => {
  $("body").html(html);
};

$(() => {
  render(genPair({ name: "PATH", value: "/usr/local/bin" }));
});
