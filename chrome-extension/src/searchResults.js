(async () => {
  function noteBlock(note) {
    return `<a class="noteCard" target="_blank" href="${note.link}">
      <h1 class="noteCardTitle">${note.title}</h1>
      <p class="noteCardSnippet">${note.snippet}</p>
    </a>`;
  }

  function asHtml(text) {
    const lines = text.split('\n').filter((el) => el !== '');
    return lines.join('<br>');
  }

  async function renderResults(results) {
    const notesBlock = document.getElementById('notesCards');
    notesBlock.innerHTML = '';

    const data = await results.json();
    if (!data || !data.entries) {
      return;
    }

    data.entries.forEach((entry) => {
      notesBlock.insertAdjacentHTML(
        'afterbegin',
        noteBlock({
          link: `vscode://file${entry.file}`,
          title: entry.title,
          snippet: asHtml(entry.slug || ''),
        }),
      );
    });
  }

  function renderError(error) {
    const notesBlock = document.getElementById('notesCards');
    notesBlock.insertAdjacentHTML(
      'afterbegin',
      noteBlock({
        link: '',
        date: '',
        title: 'ERROR',
        snippet: error,
      }),
    );
  }

  async function getSearchQuery() {
    const tabs = await chrome.tabs.query({ active: true, currentWindow: true });
    const url = new URL(tabs[0].url);
    const params = new URLSearchParams(url.search);

    return params.get('q');
  }

  const q = await getSearchQuery();
  try {
    const result = await fetch(`http://localhost:3000/?query=${q}`);
    renderResults(result);
  } catch (err) {
    renderError(err);
  }
})();
