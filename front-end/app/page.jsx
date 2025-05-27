


export default async function Root() {

  try {

    const response = await fetch("http://localhost:8080")

    const data = await response.json()

    if (data.status) {
      return (

        <div className="">
          <h1> Wellcome, sohachimi </h1>
        </div>
      )
    } else {
      return (

        <div className="">
          <h1> you have to login ! </h1>
        </div>

      )
    }

  } catch (error) {
    console.error(error);
  }

}
