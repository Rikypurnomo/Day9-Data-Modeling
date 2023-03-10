package main

import (
	"Routing/connection"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/gorilla/mux"
)

var Data = map[string]interface{}{
	"Title":   "Personal web",
	"IsLogin": false,
}

// deklarasi variable data project tipe struct
type dataProject struct {
	Id           int
	ProjectName  string
	StartDate    time.Time
	EndDate      time.Time
	Description  string
	Technologies []string
	Duration     string
}

// variable untuk menampung dataproject array of objeck
var Projects = []dataProject{}

func main() {
	route := mux.NewRouter()

	connection.DatabaseConnect()

	route.PathPrefix("/public/").Handler(http.StripPrefix("/public/",
		http.FileServer(http.Dir("./public"))))

	route.HandleFunc("/home", index).Methods("GET")
	// route.HandleFunc("/project", projectForm).Methods("GET")
	// route.HandleFunc("/project/{id}", projectDetail).Methods("GET")
	// route.HandleFunc("/project", projectAdd).Methods("POST")
	// route.HandleFunc("/contact", contactMe).Methods("GET")
	// route.HandleFunc("/delete/{id}", deleteProject).Methods("GET")
	// route.HandleFunc("/edit/{id}", editProject).Methods("GET")
	// route.HandleFunc("/editInput/{id}", editProjectInput).Methods("POST")

	port := 5000
	fmt.Println("Server is running on port", port)
	http.ListenAndServe("localhost:5000", route)
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// var tmpl, err = template.ParseFiles("views/index.html")

	var tmpl, err = template.ParseFiles("views/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("message :" + err.Error()))
		return
	}
	rows, _ := connection.Conn.Query(context.Background(), "SELECT * FROM tb_project")

	var result []dataProject
	for rows.Next() {
		var each = dataProject{}

		var err = rows.Scan(&each.Id, &each.ProjectName, &each.StartDate, &each.EndDate, &each.Description, &each.Technologies)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		each.Duration = selisihDate(each.StartDate, each.EndDate)
		result = append(result, each)
	}

	respData := map[string]interface{}{
		"Data":     Data,
		"Projects": result,
	}
	fmt.Println(result)
	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, respData)
}

func selisihDate(start time.Time, end time.Time) string {

	distance := end.Sub(start)

	// Menghitung durasi
	//pengkondisian
	var duration string
	year := int(distance.Hours() / (12 * 30 * 24))
	if year != 0 {
		duration = strconv.Itoa(year) + " tahun"
	} else {
		month := int(distance.Hours() / (30 * 24))
		if month != 0 {
			duration = strconv.Itoa(month) + " bulan"
		} else {
			week := int(distance.Hours() / (7 * 24))
			if week != 0 {
				duration = strconv.Itoa(week) + " minggu"
			} else {
				day := int(distance.Hours() / (24))
				if day != 0 {
					duration = strconv.Itoa(day) + " hari"
				}
			}
		}
	}
	return duration
}

// func projectForm(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")

// 	var tmpl, err = template.ParseFiles("views/myproject.html")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("message :" + err.Error()))
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	tmpl.Execute(w, Data)
// }

// func projectDetail(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	var tmpl, err = template.ParseFiles("views/project-detail.html")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("message :" + err.Error()))
// 		return
// 	}

// 	projectDetail := dataProject{}
// 	// perulangan
// 	for index, data := range Projects {
// 		if index == id {

// 			newStarDate, _ := time.Parse("2006-01-02", data.StartDate)
// 			newEndDate, _ := time.Parse("2006-01-02", data.EndDate)

// 			projectDetail = dataProject{
// 				// masukin data
// 				Id:          id,
// 				ProjectName: data.ProjectName,
// 				// star date nya diubah jadi string dngn format "02 Jan 2006"
// 				StartDate:    newStarDate.Format("02 Jan 2006"),
// 				EndDate:      newEndDate.Format("02 Jan 2006"),
// 				Description:  data.Description,
// 				Technologies: data.Technologies,
// 				Duration:     data.Duration,
// 			}
// 		}
// 	}

// 	resp := map[string]interface{}{
// 		"Data":          Data,
// 		"ProjectDetail": projectDetail,
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	tmpl.Execute(w, resp)
// }

// func projectAdd(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// untuk mengambil data dari input
// 	project_name := r.PostForm.Get("name")
// 	start_date := r.PostForm.Get("start-date")
// 	end_date := r.PostForm.Get("end-date")
// 	description := r.PostForm.Get("message")
// 	technologies := r.Form["project-tech"]

// 	start_date_time, _ := time.Parse("2006-01-02", start_date)
// 	end_date_time, _ := time.Parse("2006-01-02", end_date)

// 	selisih_date := end_date_time.Sub(start_date_time)

// 	// untuk menghitung duration
// 	var duration string
// 	year := int(selisih_date.Hours() / (12 * 30 * 24))
// 	if year != 0 {
// 		duration = strconv.Itoa(year) + " tahun"
// 	} else {
// 		month := int(selisih_date.Hours() / (30 * 24))
// 		if month != 0 {
// 			duration = strconv.Itoa(month) + " bulan"
// 		} else {
// 			week := int(selisih_date.Hours() / (7 * 24))
// 			if week != 0 {
// 				duration = strconv.Itoa(week) + " minggu"
// 			} else {
// 				day := int(selisih_date.Hours() / (24))
// 				if day != 0 {
// 					duration = strconv.Itoa(day) + " hari"
// 				}
// 			}
// 		}
// 	}

// 	// untuk memasukan data yg dipanggil dari inputform ke variable data project
// 	var newProject = dataProject{
// 		ProjectName:  project_name,
// 		StartDate:    start_date,
// 		EndDate:      end_date,
// 		Technologies: technologies,
// 		Description:  description,
// 		Duration:     duration,
// 	}

// 	// append untuk menambahkan data di project ke newproject
// 	Projects = append(Projects, newProject)
// 	fmt.Println("Name :" + r.PostForm.Get("name"))
// 	fmt.Println("Start :" + r.PostForm.Get("start-date"))
// 	fmt.Println("End :" + r.PostForm.Get("end-date"))
// 	fmt.Println("Description :" + r.PostForm.Get("message"))
// 	// r.Form ambil data lebih dari 1
// 	fmt.Println("Tech Stack :", r.Form["project-tech"])

// 	http.Redirect(w, r, "/home", http.StatusMovedPermanently)
// }

// func contactMe(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")

// 	var tmpl, err = template.ParseFiles("views/contact-form.html")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("message :" + err.Error()))
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	tmpl.Execute(w, Data)
// }

// func deleteProject(w http.ResponseWriter, r *http.Request) {
// 	// ngambil id pake ini
// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	Projects = append(Projects[:id], Projects[id+1:]...)

// 	http.Redirect(w, r, "/home", http.StatusFound)
// }

// func editProject(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")

// 	var tmpl, err = template.ParseFiles("views/edit.html")
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("message :" + err.Error()))
// 		return
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	ProjectDetail := dataProject{}
// 	// perulangan
// 	for index, data := range Projects {
// 		// ==2x eta equal
// 		if index == id {

// 			ProjectDetail = dataProject{
// 				// masukin data
// 				Id:           id,
// 				ProjectName:  data.ProjectName,
// 				StartDate:    data.StartDate,
// 				EndDate:      data.EndDate,
// 				Description:  data.Description,
// 				Technologies: data.Technologies,
// 				Duration:     data.Duration,
// 			}
// 		}
// 	}

// 	respData := map[string]interface{}{
// 		"Data":          Data,
// 		"ProjectDetail": ProjectDetail,
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	tmpl.Execute(w, respData)
// }

// func editProjectInput(w http.ResponseWriter, r *http.Request) {

// 	err := r.ParseForm()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	projectName := r.PostForm.Get("name")
// 	startDate := r.PostForm.Get("start-date")
// 	endDate := r.PostForm.Get("end-date")
// 	description := r.PostForm.Get("message")
// 	techStack := r.Form["project-tech"]

// 	// Menghitung durasi
// 	// Parsing string ke time

// 	// Start Date
// 	startDateTime, _ := time.Parse("2006-01-02", startDate)

// 	// End Date
// 	endDateTime, _ := time.Parse("2006-01-02", endDate)

// 	// Perbedaan nya berupa : jam menit detik
// 	distance := endDateTime.Sub(startDateTime)

// 	// Menghitung durasi
// 	//pengkondisian
// 	var duration string
// 	year := int(distance.Hours() / (12 * 30 * 24))
// 	if year != 0 {
// 		duration = strconv.Itoa(year) + " tahun"
// 	} else {
// 		month := int(distance.Hours() / (30 * 24))
// 		if month != 0 {
// 			duration = strconv.Itoa(month) + " bulan"
// 		} else {
// 			week := int(distance.Hours() / (7 * 24))
// 			if week != 0 {
// 				duration = strconv.Itoa(week) + " minggu"
// 			} else {
// 				day := int(distance.Hours() / (24))
// 				if day != 0 {
// 					duration = strconv.Itoa(day) + " hari"
// 				}
// 			}
// 		}
// 	}

// 	var newProject = dataProject{
// 		ProjectName:  projectName,
// 		StartDate:    startDate,
// 		EndDate:      endDate,
// 		Description:  description,
// 		Technologies: techStack,
// 		Duration:     duration,
// 	}

// 	id, _ := strconv.Atoi(mux.Vars(r)["id"])

// 	Projects[id] = newProject

// 	http.Redirect(w, r, "/home", http.StatusMovedPermanently)

// }
