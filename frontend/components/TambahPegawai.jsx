import axios from 'axios'
import React, { Component } from 'react'
import styles from '../pages/index.module.css'
import Link from 'next/link'

export default class TambahPegawai extends Component {
  constructor(props) {
    super(props)
    this.state = {
      berhasil_tambah: null,
      data: {
        nama: "lorem",
        alamat: "ipsum",
        telepon: "dayeuh kolot"
      }
    }
    this.berhasil_tambah = this.berhasil_tambah.bind(this)
  }

  berhasil_tambah = () => {
    if (this.state.berhasil_tambah == false) {
      return (
        <div className={styles.container} style={{ marginTop: "1rem" }}>
          <h2 className="text-red-500">Tidak berhasil menambahkan pegawai</h2>
        </div>
      )
    } else if (this.state.berhasil_tambah == true) {
      return (
        <div className={styles.container} style={{ marginTop: "1rem" }}>
          <h2 className="text-green-500">Berhasil Menambahkan Pegawai</h2>
          <p>Data Pegawai:</p>
          <p>Nama:  {this.state.data.nama}</p>
          <p>Alamat:  {this.state.data.alamat}</p>
          <p>Telepon:  {this.state.data.telepon}</p>
        </div>
      )
    }
  }

  postPegawai = (event) => {
    event.preventDefault()
    let nama_inp = event.target.nama.value
    let alamat_inp = event.target.alamat.value
    let telepon_inp = "+62"+`${event.target.telepon.value}`

    var bodyFormData = new FormData()

    bodyFormData.append('nama', nama_inp)
    bodyFormData.append('alamat', alamat_inp)
    bodyFormData.append('telepon', telepon_inp)


    axios({
      method: 'post',
      url : '/api/pegawai',
      data: bodyFormData,
      headers : { "Content-Type": "multipart/form-data" }
    }).then(res => {
      let data_rest = res.data.data
      let status = res.status
      if (status == 200) {
        this.setState({data : {
          nama: data_rest.nama,
          alamat: data_rest.alamat,
          telepon: data_rest.telepon
        }, berhasil_tambah: true})
      } else {
        this.setState({berhasil_tambah : false})
      }
     
    }).catch(e => console.log(e))
  }

  render() {
    return (
      <div>
        <div className={styles.container}>
          <h1 className>Tambah Pegawai</h1>

          <form onSubmit={this.postPegawai} className="mt-4 grid grid-flow-row">
            <label htmlFor="nama">
              Nama:
              <input type="text" name="nama" id="nama" placeholder="Masukkan nama lengkap" required/>
            </label>
            <label htmlFor="telepon">
              Telepon:
              <div className="flex">
                <p className="mt-auto mb-auto m-1">+62</p>
                <input type="number" name="telepon" id="telepon" placeholder="8572830777" required/>
              </div>
            </label>
            <label htmlFor="alamat" className="grid grid-flow-row">
              Alamat:
              <textarea name="alamat" id="alamat" rows="10" placeholder="Masukkan Nama Jalan RT/RW, Keluarahan, Kecamatan, Provinsi, Kode pos" ></textarea>
            </label>

            <button type="submit" className="mt-5 w-1/2 mr-auto ml-auto">Submit</button>
          </form>
        </div>


        {this.berhasil_tambah()}
      </div>
    )
  }
}
