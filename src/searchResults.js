(async () => {
  function noteBlock(note) {
    return `<a class="noteCard" target="_blank" href="${note.link}">
      <h1 class="noteCardTitle">${note.title}</h1>
      <p class="noteCardDate">${note.date}</p>
      <p class="noteCardSnippet">${note.snippet}</p>
    </a>`;
  }

  async function getSearchQuery() {
    const tabs = await chrome.tabs.query({ active: true, currentWindow: true });
    const url = new URL(tabs[0].url);
    const params = new URLSearchParams(url.search);

    return params.get('q');
  }

  const q = await getSearchQuery();
  console.log(q);

  const notesBlock = document.getElementById('notesCards');
  notesBlock.insertAdjacentHTML(
    'afterbegin',
    noteBlock({
      link: '',
      date: '2022-08-20',
      title: 'Example',
      snippet: 'Some text is hehe',
    }),
  );
  notesBlock.insertAdjacentHTML(
    'afterbegin',
    noteBlock({
      link: '',
      date: '2022-08-20',
      title: 'Example',
      snippet: 'Some text is hehe',
    }),
  );
  notesBlock.insertAdjacentHTML(
    'afterbegin',
    noteBlock({
      link: '',
      date: '2022-08-20',
      title: 'Example',
      snippet: 'Some text is hehe',
    }),
  );
})();
