package main

import (
	"time"
)

// Patient represents a patient in the healthcare system
type Patient struct {
	ID            string
	Name          string
	Birthday      time.Time
	MedicalRecord *MedicalRecord
}

// MedicalRecord represents a patient's medical record in the healthcare system
type MedicalRecord struct {
	PatientID   string
	Conditions  []string
	Medications []string
	Procedures  []string
}

// AddCondition adds a new condition to the patient's medical record
func (p *Patient) AddCondition(condition string) {
	p.MedicalRecord.Conditions = append(p.MedicalRecord.Conditions, condition)
}

// AddMedication adds a new medication to the patient's medical record
func (p *Patient) AddMedication(medication string) {
	p.MedicalRecord.Medications = append(p.MedicalRecord.Medications, medication)
}

// AddProcedure adds a new procedure to the patient's medical record
func (p *Patient) AddProcedure(procedure string) {
	p.MedicalRecord.Procedures = append(p.MedicalRecord.Procedures, procedure)
}

func main() {
	// Create a new patient
	patient := &Patient{
		ID:       "123",
		Name:     "Alice",
		Birthday: time.Now(),
		MedicalRecord: &MedicalRecord{
			PatientID:   "123",
			Conditions:  []string{},
			Medications: []string{},
			Procedures:  []string{},
		},
	}

	// Add a new condition to the patient's medical record
	patient.AddCondition("Flu")

	// Add a new medication to the patient's medical record
	patient.AddMedication("Aspirin")

	// Add a new procedure to the patient's medical record
	patient.AddProcedure("Flu shot")
}
