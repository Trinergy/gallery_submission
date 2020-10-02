import React from 'react';

class ImageList extends React.Component {
  constructor() {
    super();
    this.state = {
      images: []
    };
  };

  componentDidMount() {
    // TODO GET API
    this.setState({
      images: [
        { id: 1, url: "https://picsum.photos/200/300", updated_at: "9:00", created_at: "1:00", flagged: false },
        { id: 2, url: "https://picsum.photos/200/300", updated_at: "9:00", created_at: "1:00", flagged: true },
        { id: 3, url: "https://picsum.photos/200/300", updated_at: "9:00", created_at: "1:00", flagged: false },
        { id: 4, url: "https://picsum.photos/200/300", updated_at: "9:00", created_at: "1:00", flagged: false },
      ]
    })
  }

  handleClick(id){
    console.log("hello",id)
    // TODO POST API TO UPDATE FLAGGED - ON SUCCESS, SETSTATE - ON FAIL, RENDER ERROR
  }


  render() {
    const imageList = this.state.images.map((i) => <li key={i.id}>
      <img src={i.url}></img> <a onClick={() => this.handleClick(i.id)} class="button is-primary"> Tag</a>
    </li>);
    return (
      <div>
        {imageList}
      </div>
    );
  }
}

export default ImageList;